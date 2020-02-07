package ktrans

import (
	"fmt"

	env "github.com/afeldman/go-util/env"
	str_util "github.com/afeldman/go-util/string"

	"os"
	"path/filepath"
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

func SearchForKtrans() (error, string) {

	for _, path := range env.GetEnv("PATH") {
		err := filepath.Walk(path, findFile)
		if err != nil {
			return err, ""
		}
	}

	if !str_util.StringEmpty(ktranspath) {
		return nil, ktranspath
	}

	return nil, ""
}
