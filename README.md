# STOML - Simple TOML parser for Linux Shell

## What is it
STOML is a simple TOML file parser for Linux Shell (bash, ksh, csh, etc).
The purpose is to easily parse INI files.
TOML is a generalized INI file, so a simplified TOML parser can get this done.

## How to get it
Check the [releases](https://github.com/freshautomations/stoml/releases) page.

## How to build it from source
```cgo
GO111MODULE=on go get github.com/freshautomations/stoml.git
```

## How to use
To get a `key` from a file called `filename` and export it as a Bash environment variable, run
```bash
export MYVALUE=`stoml filename key`
```
TOML can also have sections. They are referred to with dotted paths:
```bash
export MYSECTIONVALUE=`stoml filename mysection.key`
```

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
#Export as "[1 2 3]"
```
Note: floats that have 9 zeros after the dot (epsilon = 10^-9) are presented as integers.
Viper's toml converter somehow converts big integers to float and presents them in the exponential format. This way we work around that.

## Cool side-effects
The application will try to figure out what kind of file was it fed.
`*.toml` files will be parsed as TOML but as a side-effect of the viper library,
the application can also parse JSON files (`*.json`) and YAML (`*.yml`) files.

By default, any other file types (for example INI) are treated as TOML files.

This side-effect is used during the release of the application to parse incoming JSON responses from the GitHub API.
It's not a main feature though.

