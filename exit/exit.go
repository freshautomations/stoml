package exit

import (
	"fmt"
	"os"
	"strings"
)

var Quiet bool
var Strict bool

func Fail(err error) {
	if !Quiet {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(1)
}

func Succeed(result string) {
	if Strict && strings.TrimSpace(result) == "" {
		if !Quiet {
			fmt.Fprintln(os.Stderr, "missing or empty field")
		}
		os.Exit(1)
	}
	fmt.Print(result)
	os.Exit(0)
}
