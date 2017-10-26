//https://stackoverflow.com/questions/35934298/golang-how-to-redirect-to-a-url/35934496
//https://stackoverflow.com/questions/28793619/golang-what-to-use-http-servefile-or-http-fileserver
//"html/template" use this when using first example 
//	"log" use this when using first example 
package main

import (
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
