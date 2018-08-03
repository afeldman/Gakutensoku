package ktrans

import (
	"fmt"

	env "github.com/afeldman/go-util/env"

	"os"
	"path/filepath"
	"unicode"
)

var ktranspath string

func findFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// get absolute path of the folder that we are searching
	absolute, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	matched, err := filepath.Match("ktrans.exe", info.Name())
	if err != nil {
		fmt.Println(err)
	}

	if matched {
		ktranspath = absolute
	}
	return nil
}

func SearchForKtrans() string {

	for _, path := range env.GetEnv("PATH") {
		err := filepath.Walk(path, findFile)
		if err != nil {
			println("Error", err)
		}
	}

	if !strempty(ktranspath) {
		return ktranspath
	} else {
		return ""
	}

}

// Strempty checks whether string contains only whitespace or not
func strempty(s string) bool {
	if len(s) == 0 {
		return true
	}

	r := []rune(s)
	l := len(r)

	for l > 0 {
		l--
		if !unicode.IsSpace(r[l]) {
			return false
		}
	}

	return true
}
