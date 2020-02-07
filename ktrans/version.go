package ktrans

import (
	"fmt"
	"regexp"
)

func Version() string {
	output := help()
	re, err := regexp.Compile(`KTRANS (V\d+.\d+-\d+),`)
	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
	}

	var ver = re.FindStringSubmatch(output)

	return ver[1]
}
