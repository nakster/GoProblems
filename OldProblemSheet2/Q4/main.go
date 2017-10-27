//https://stackoverflow.com/questions/35934298/golang-how-to-redirect-to-a-url/35934496
package main

import (
	"net/http"
	"html/template"
	"log"
)
type messagStruct struct{

	Message string

}
//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	  http.ServeFile(w, r, "index.html") 
}
//can name it anything you want i.e handlerFunc 
func handleGuess(w http.ResponseWriter, r *http.Request) {
	//create a new variable and assign the struct and string value
	 mess := &messagStruct{
		 Message: "Guess a number between 1 and 20"}
	
	 //this parses the html file 
	t, err := template.ParseFiles("guess.tmpl") 
	 //error handling 
    if err != nil { 
  	  log.Print("error ", err) 
  	}
	//excute sinfo struct 
   err = t.Execute(w, mess) 
    if err != nil { 
  	  log.Print("error ", err) 
  	}
}

func main() {
	///////////open localhost on port 8080
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handleGuess)
	http.ListenAndServe(":8080", nil)
}
