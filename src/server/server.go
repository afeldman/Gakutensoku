package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Start_Mux() {

	r := mux.NewRouter()
	r.HandleFunc("/", ReceiveFile).Methods("POST")
	//r.HandleFunc("/articles", handler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    *server + ":" + *port,
	}

	log.Fatal(srv.ListenAndServe())

}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer

	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	// Copy the file data to my buffer
	io.Copy(&Buf, file)

	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	contents := Buf.String()
	fmt.Println(contents)

	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	Buf.Reset()

	// do something else
	// etc write header
	return
}
