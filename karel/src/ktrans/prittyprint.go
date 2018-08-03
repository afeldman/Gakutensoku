package ktrans

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printCommand(cmd *exec.Cmd) {
	if debug {
		fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
	}
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if (len(outs) > 0) && (debug) {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
