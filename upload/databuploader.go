package routes

import (
	"errors"
	"io"

	log "github.com/sirupsen/logrus"

	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/*
curl -X POST -i -H "Accept: * /*" \
	-H "Accept-Encoding: gzip, deflate, br" \
	-H "Connection: keep-alive" \
	-H "Content-Type: multipart/form-data" \
	-F file="<data>" http://localhost:2611/file
*/
func UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Debugln("upload a filefunction")
	file, handler, err := r.FormFile("file")

	fileName := handler.Filename
	log.Debugln("file name is: " + fileName)

	fileExtension := filepath.Ext(fileName)
	if !strings.EqualFold(fileExtension, ".karel") {
		log.Errorln("file extension " + fileExtension + " not karel type")
		_, _ = io.WriteString(w, "file extension "+fileExtension+" not karel type")
		return
	}

	upload_filepath := filepath.Join(os.Getenv("FILE_STORAGE"), "uploads")
	log.Println("upload path is: " + upload_filepath)

	log.Debugln("check if upload path exists")
	if _, err := os.Stat(upload_filepath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(upload_filepath, os.ModePerm)
		if err != nil {
			log.Errorln(err)
		}
	}

	log.Debugln("open file to copy data in")
	f, err := os.OpenFile(filepath.Join(upload_filepath, fileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
	_, _ = io.Copy(f, file)

	// f is copied now work on the karel system.
	log.Debugln("open the file to check for kpc file. if not kpc set delete the file")

}
