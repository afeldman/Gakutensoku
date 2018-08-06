package ktrans

import "os/exec"

func help() string {
	cmd := exec.Command("ktrans")
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)

	return string(output)
}
