package upload

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseContentRange(contentRange string) (totalSize int64, partFrom int64, partTo int64) {
	contentRange = strings.Replace(contentRange, "bytes ", "", -1)
	fromTo := strings.Split(contentRange, "/")[0]
	totalSize, err := strconv.ParseInt(strings.Split(contentRange, "/")[1], 10, 64)
	checkError(err)

	splitted := strings.Split(fromTo, "-")

	partFrom, err = strconv.ParseInt(splitted[0], 10, 64)
	checkError(err)
	partTo, err = strconv.ParseInt(splitted[1], 10, 64)
	checkError(err)

	return totalSize, partFrom, partTo
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}

	return false
}

func ensureDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
}
