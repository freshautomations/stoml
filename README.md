# STOML - Simple TOML parser for the Linux Shell

## What is it
STOML is a simple TOML file parser for the Linux Shell (bash, ksh, csh, etc).
The purpose is to easily parse TOML and INI files.
TOML is a generalized INI file, so a simplified TOML parser can get this done.
Currently, INI files are always parsed as strings. TOML files are parsed by type as described below.

Tom's Obvious, Minimal Language is defined [here](https://github.com/toml-lang/toml).

## How to get it
Check the [releases](https://github.com/freshautomations/stoml/releases) page.

## How to build it from source
Install [Golang](https://golang.org/doc/install), then run:
```cgo
GO111MODULE=on go get github.com/freshautomations/stoml.git
```

## How to use
To get a `key` from a file called `filename` and export it as a Bash environment variable, run
```bash
export MYVALUE=`stoml filename key`
```
TOML can also have tables (sections). They are referred to with dotted paths:
```bash
export MYSECTIONVALUE=`stoml filename mysection.key`
```
Error messages can be suppressed using `-q`. This is useful when running in a script. 
## Caveats
This is a `simple` implementation of parsing TOML.
This means that it will work well with atomic types in the configuration, like string, int or boolean.
It also means that more complex types like floats, lists or maps will be translated by Go as it sees fit.

Examples for output:
```toml
myint = +4
#Export as "4"

myfakeint = "+4"
#Export as "+4"

mystring = "hello hi g'day howdy"
#Export as "hello hi g'day howdy"

mylist = [1,2,3]
#Export as "1 2 3"
```
Note: floats that have 9 zeros after the dot (epsilon = 10^-9) are presented as integers.
Viper's toml converter somehow converts big integers to float and presents them in the exponential format. This way we work around that.

## Bugs
* INI-files cannot read the list of section keys. (TOML files can.) Not sure of the reason yet.
* Empty sections are invisible and not returned when requesting the list of sections. This is an implementation detail in the Viper library. Issue [here](https://github.com/spf13/viper/issues/1131).
* Dotted keys have to be quoted for now. The [TOML specification](https://toml.io/en/v1.0.0) says otherwise, but Viper seems to disagree. Issue [here](https://github.com/freshautomations/stoml/issues/2).

I'll keep an eye out for solutions to these issues.

## INI-file parsing
INI-files are more relaxed TOML files. There are minor differences in how they are parsed:
* INI files only support strings. Everything will be read as a string, as is. (TOML will make arrays look nice and Shell-friendly.)
* There is an unmarked "[default]" section in INI files (the root table), all top-level entries will be available through this default table.

Check the tests for more details.

## Cool side-effects
The application will try to figure out what kind of file it was fed.
`*.toml` files will be parsed as TOML but as a side-effect of the viper library,
the application can also parse JSON files (`*.json`) and YAML (`*.yml`) files.

By default, any other file types are treated as INI files.

This side-effect is used during the release of the application to parse incoming JSON responses from the GitHub API.
It's not a main feature though.
