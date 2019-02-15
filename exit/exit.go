package exit

import (
	"fmt"
	"os"
)

func Fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func Succeed(result string) {
	fmt.Print(result)
	os.Exit(0)
}
