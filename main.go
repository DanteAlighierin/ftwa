package main

import {
	"net/http"
	"fmt"
}


func uploadFile(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Uploading File")
}


func setupRotes(){
	http.HandleFunc("/upload", uploadFile")
	http.ListenAndServe("8000",nil)
}


func main(){
	fmt.Println("Go File Upload")
	setupRoutes()
}
