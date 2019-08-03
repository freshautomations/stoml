package cmd

import (
	"github.com/freshautomations/stoml/defaults"
	"github.com/freshautomations/stoml/exit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sort"
	"strings"
)

const epsilon = 1e-9 // Margin of error

func CheckArgs(cmd *cobra.Command, args []string) error {
	validateArgs := cobra.ExactArgs(2)
	if err := validateArgs(cmd, args); err != nil {
		return err
	}

	fileName := args[0]
	_, err := os.Stat(fileName)
	return err
}

func RunRoot(cmd *cobra.Command, args []string) (output string, err error) {
	fileName := args[0]
	key := args[1]

	viper.SetConfigFile(fileName)
	err = viper.ReadInConfig()
	if err != nil {
		if _, IsUnsupportedExtension := err.(viper.UnsupportedConfigError); IsUnsupportedExtension {
			viper.SetConfigType("toml")
			err = viper.ReadInConfig()
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	switch viper.Get(key).(type) {
	// If a section was requested, return the keys from that section
	case map[string]interface{}:
		var r []string
		for k, _ := range viper.GetStringMapString(key) {
			r = append(r,k)
		}
		sort.Strings(r)
		output = strings.Join(r, " ")
	// Return list of strings as result
	default:
		// Return all section names and root section keys if "." is provided
		if key == "." {
			var r []string
			for k, _ := range viper.AllSettings() {
				r = append(r,k)
			}
			sort.Strings(r)
			output = strings.Join(r, " ")
		} else {
			output = strings.Join(viper.GetStringSlice(key), " ")
		}
	}
	return
}

func runRootWrapper(cmd *cobra.Command, args []string) {
	if result, err := RunRoot(cmd, args); err != nil {
		exit.Fail(err)
	} else {
		exit.Succeed(result)
	}
}

func Execute() error {
	var rootCmd = &cobra.Command{
		Version: defaults.Version,
		Use:     "stoml",
		Short:   "STOML - simple toml parser for Shell",
		Long: `A simplified TOML (also known as a more formal INI) parser for the Linux Shell.
Source and documentation is available at https://github.com/freshautomations/stoml`,
		Args: CheckArgs,
		Run:  runRootWrapper,
	}
	rootCmd.Use = "stoml <filename> <key>"

	return rootCmd.Execute()
}
