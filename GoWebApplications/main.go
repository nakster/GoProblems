package main

import (
	//"fmt"
	"net/http"
)

//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}


func main() {
	///////////open localhost on port 8080
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
