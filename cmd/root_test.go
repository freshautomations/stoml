package cmd

import (
	"github.com/freshautomations/stoml/defaults"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckArgs(t *testing.T) {
	cmd := &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Version: defaults.Version,
	}

	assert.NotNil(t, CheckArgs(cmd, []string{"../test.ini"}), "enough parameters")
	assert.NotNil(t, CheckArgs(cmd, []string{"notexist.ini", "master_of_the_universe"}), "file found")
	assert.Nil(t, CheckArgs(cmd, []string{"../test.ini", "master_of_the_universe"}), "parameter check")
}

func TestRunRoot(t *testing.T) {
	cmd := &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Version: defaults.Version,
	}

	var result interface{}
	var err error

	// Root section
	result, err = RunRoot(cmd, []string{"../test.ini", "master_of_the_universe"})
	assert.Equal(t, "false", result, "unexpected result")
	assert.Nil(t, err, "unexpected error")

	// Section
	result, err = RunRoot(cmd, []string{"../test.ini", "district9.name"})
	assert.Equal(t, "Wikus", result, "unexpected result")
	assert.Nil(t, err, "unexpected error")

	// Case-insensitive
	result, err = RunRoot(cmd, []string{"../test.ini", "district9.eta"})
	assert.Equal(t, "3", result, "unexpected result")
	assert.Nil(t, err, "unexpected error")

	// No key
	result, err = RunRoot(cmd, []string{"../test.ini", "blur2.song3"})
	assert.Empty(t, result, "unexpected result")
	assert.Nil(t, err, "unexpected error")

	// No section
	result, err = RunRoot(cmd, []string{"../test.ini", "blur3.song3"})
	assert.Empty(t, result, "unexpected result")
	assert.Nil(t, err, "unexpected error")

	// invalid input
	result, err = RunRoot(cmd, []string{"../test.ini", "#invalid"})
	assert.Empty(t, result, "unexpected result")
	assert.Nil(t, err, "unexpected error")

	// Todo: FIX long int from json test
	//result, err = RunRoot(cmd, []string{"../test.json", "id"})
	//assert.Equal(t, "15576104", result, "unexpected result")
	//assert.Nil(t, err, "unexpected error")

}
