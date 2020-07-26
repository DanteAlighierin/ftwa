package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	///1. parse input, , type multipart/form-data
	r.ParseMultipartForm(10 << 20)

	//2. retrieve file from posted form data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error")
		fmt.Println("err")
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Filename)
	fmt.Printf("MIME type: %+v\n", handler.Header)

	//3. write temp file to server
	tempFile, err := ioutil.TempFile("./temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	//4. return whether or not  this has ben succesful
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func main() {
	fmt.Println("Go File Upload")
	fmt.Println("The app is running at http://localhost:8080/static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	setupRoutes()

}
