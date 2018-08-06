package ktrans

import (
	"os/exec"
	"strings"

	string "github.com/afeldman/go-util/string"
)

var debug = true

type Ktrans struct {
	Path, Out, In, Ver, Conf string
	Param                    map[string]bool
}

func Init() *Ktrans {

	this := new(Ktrans)

	this.Param = make(map[string]bool)
	this.Param["l"] = false
	this.Param["r"] = false
	this.Param["d"] = false
	this.Param["p"] = false

	this.Conf = ""
	this.In = ""
	this.Out = ""
	this.Path = SearchForKtrans()
	this.Ver = ""

	return this
}

func (this *Ktrans) Cmd() []byte {

	var cmd_string []string
	cmd_string = append(cmd_string, this.Path)

	for k, v := range this.Param {
		if v {
			cmd_string = append(cmd_string, "/"+k)
		}
	}

	if !string.StringEmpty(this.Ver) {
		cmd_string = append(cmd_string, "/ver"+this.Ver)
	}
	if !string.StringEmpty(this.In) {
		cmd_string = append(cmd_string, this.In)
	}
	if !string.StringEmpty(this.Out) {
		cmd_string = append(cmd_string, this.Out)
	}
	if !string.StringEmpty(this.Conf) {
		cmd_string = append(cmd_string, "/config", this.Conf)
	}

	cmd := exec.Command(strings.Join(cmd_string, " "))
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)

	return output
}
