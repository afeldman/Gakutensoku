package ktrans

import (
	"os/exec"
	"strings"

	"log"

	print "github.com/afeldman/go-util/print"
	str_util "github.com/afeldman/go-util/string"
)

type Ktrans struct {
	PathToKtrans      string `json:"ktrans" yaml:"ktrans"`
	Output            string `json:"output,omitempty" yaml:"output,omitempty"`
	Input             string `json:"input" yaml:"input"`
	Version           string `json:"version" yaml:"version"`
	ConfigurationFile string `json:"robot,omitempty" yaml:"robot,omitempty"`

	L bool `json:"list,omitempty" yaml:"list,omitempty"`
	R bool `json:"routine,omitempty" yaml:"routine,omitempty"`
	Q bool `json:"pause,omitempty" yaml:"pause,omitempty"`
	D bool `json:"display,omitempty" yaml:"display,omitempty"`
}

func Init() *Ktrans {

	this := new(Ktrans)

	this.L = false
	this.R = false
	this.D = false
	this.Q = false

	err, ktranspath := SearchForKtrans()
	if err != nil {
		log.Println(err)
	}
	if !str_util.StringEmpty(ktranspath) {
		log.Println("No KTrans found in $PATH")
	}

	this.ConfigurationFile = ""
	this.Input = ""
	this.Output = ""
	this.PathToKtrans = ktranspath
	this.Version = ""

	return this
}

//Cmd This method builds the
func (this *Ktrans) Cmd() []byte {

	var cmdstring []string
	cmdstring = append(cmdstring, this.PathToKtrans)

	if this.R {
		cmdstring = append(cmdstring, "/r")
	}
	if this.D {
		cmdstring = append(cmdstring, "/d")
	}
	if this.L {
		cmdstring = append(cmdstring, "/l")
	}
	if this.Q {
		cmdstring = append(cmdstring, "/p")
	}
	if !str_util.StringEmpty(this.Version) {
		cmdstring = append(cmdstring, "/ver"+this.Version)
	}
	if !str_util.StringEmpty(this.Input) {
		cmdstring = append(cmdstring, this.Input)
	}
	if !str_util.StringEmpty(this.Output) {
		cmdstring = append(cmdstring, this.Output)
	}
	if !str_util.StringEmpty(this.ConfigurationFile) {
		cmdstring = append(cmdstring, "/config", this.ConfigurationFile)
	}

	cmd := exec.Command(strings.Join(cmdstring, " "))
	print.PrintCommand(cmd)
	output, err := cmd.CombinedOutput()
	print.PrintError(err)
	print.PrintOutput(output)

	return output
}
