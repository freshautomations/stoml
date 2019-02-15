package main

import (
	"github.com/freshautomations/stoml/cmd"
	"github.com/freshautomations/stoml/exit"
)

func main() {
	if err := cmd.Execute(); err != nil {
		exit.Fail(err)
	}
}
