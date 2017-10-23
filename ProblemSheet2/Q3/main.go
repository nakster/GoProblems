//https://stackoverflow.com/questions/35934298/golang-how-to-redirect-to-a-url/35934496
package main

import (
	"fmt"
	"net/http"
)
type pageTitle struct{

	pageT string

}
//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Guessing game")

	 t, err := template.ParseFiles("Q3.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}

}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "New Game")
}

func main() {
	///////////open localhost on port 8080
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handlerFunc2)
	http.ListenAndServe(":8080", nil)
}
