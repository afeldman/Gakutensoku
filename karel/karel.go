package karel

import (
	"fmt"
	"runtime"

	ktrans "github.com/afeldman/Gakutensoku/ktrans"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println(ktrans.Version())
		test := ktrans.Init()
		test.Cmd()
	} else {
		fm
	}
}
