package karel

import (
	"fmt"
	"runtime"

	ktrans "github.com/afeldman/Gakutensoku/ktrans"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println(ktrans.Version())
		test := ktrans.InitKTrans()
		test.Cmd()
	} else {
		fmt.Println("for the momement there is not unix version. FANUC only supports Windows")
	}
}
