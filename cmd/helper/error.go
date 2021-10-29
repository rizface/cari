package helper

import (
	"fmt"
	"os"
)

func PrintIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
