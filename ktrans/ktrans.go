package ktrans

import (
	"os/exec"
	"strings"

	string "github.com/afeldman/go-util/string"
	print "github.com/afeldman/go-util/print"
)

var debug = true

type Ktrans struct {
	PathToKtrans       string `json:"ktrans" yaml:"ktrans"`
	Output            string `json:"output,omitempty" yaml:"output,omitempty"`
	Input             string `json:"input" yaml:"input"`
	Version           string `json:"version" yaml:"version"`
	ConfigurationFile string `json:"robot,omitempty" yaml:"robot,omitempty"`

	l bool `json:"list,omitempty" yaml:"list,omitempty"`
	r bool `json:"routine,omitempty" yaml:"routine,omitempty"`
	p bool `json:"pause,omitempty" yaml:"pause,omitempty"`
	d bool `json:"display,omitempty" yaml:"display,omitempty"`
}

func KtransInit() *Ktrans {

	this := new(Ktrans)

	this.l = false
	this.r = false
	this.d = false
	this.p = false

	this.ConfigurationFile = ""
	this.Input = ""
	this.Output = ""
	this.PathToKtrans = SearchForKtrans()
	this.Version = ""

	return this
}

func (this *Ktrans) Cmd() []byte {

	var cmd_string []string
	cmd_string = append(cmd_string, this.PathToKtrans)


	if this.r {
		cmd_string = append(cmd_string, "/r")
	}
	if this.d {
		cmd_string = append(cmd_string, "/d")
	}
	if this.l {
		cmd_string = append(cmd_string, "/l")
	}
	if this.p {
		cmd_string = append(cmd_string, "/p")
	}
	if !str_util.StringEmpty(this.Version) {
		cmd_string = append(cmd_string, "/ver"+this.Version)
	}
	if !str_util.StringEmpty(this.Input) {
		cmd_string = append(cmd_string, this.In)
	}
	if !str_util.StringEmpty(this.Output) {
		cmd_string = append(cmd_string, this.Out)
	}
	if !str_util.StringEmpty(this.ConfigurationFile) {
		cmd_string = append(cmd_string, "/config", this.Conf)
	}

	cmd := exec.Command(strings.Join(cmd_string, " "))
	print.PrintCommand(cmd)
	output, err := cmd.CombinedOutput()
	print.PrintError(err)
	print.PrintOutput(output)

	return output
}
