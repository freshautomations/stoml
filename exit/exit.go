package exit

import (
	"fmt"
	"os"
)

var Quiet bool

func Fail(err error) {
	if !Quiet {
		fmt.Println(err)
	}
	os.Exit(1)
}

func Succeed(result string) {
	fmt.Print(result)
	os.Exit(0)
}
