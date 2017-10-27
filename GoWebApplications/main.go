package main

import (
	"fmt"
	"net/http"
)

//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Guessing game</h1>")
}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>New Game</h1>")
}

func main() {
	///////////open localhost on port 8080
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handlerFunc2)
	http.ListenAndServe(":8080", nil)
}
