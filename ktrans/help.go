package ktrans

import (
	"os/exec"

	print "github.com/afeldman/go-util/print"
)
func help() string {
	cmd := exec.Command("ktrans")
	print.PrintCommand(cmd)
	output, err := cmd.CombinedOutput()
	print.PrintError(err)
	print.PrintOutput(output)

	return string(output)
}
