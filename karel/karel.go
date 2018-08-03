package main

import (
	"fmt"
	ktrans "github.com/afeldman/karel/src/ktrans"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println(ktrans.Version())
		test := ktrans.Init()
		test.Cmd()
	}
}
