package helper

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(link string) {
	os := runtime.GOOS
	if os == "linux" {
		exec.Command("xdg-open", link).Start()
	} else if os == "windows" {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", link).Start()
	} else if os == "darwin" {
		exec.Command("open", link).Start()
	} else {
		fmt.Println("unsupported platform")
	}

	fmt.Println("selesai")
}
