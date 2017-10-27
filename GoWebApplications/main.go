//https://stackoverflow.com/questions/35934298/golang-how-to-redirect-to-a-url/35934496
package main

import (
	"net/http"
	"html/template"
	"log"
	"strconv"
	"time"
	"math/rand"
	"fmt"
)
type messagStruct struct{

	Message string
	Guess int
	NewMessage string
	Winner bool

}
//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	  http.ServeFile(w, r, "index.html") 
}
//can name it anything you want i.e handlerFunc 
func handleGuess(w http.ResponseWriter, r *http.Request) {

	randomNum := getRandumNum(1, 20)
	//set a new cookie if doesnt exists 
	cookie, err := r.Cookie("target"); 

	if err != nil {
		cookie = &http.Cookie{Name: "target", Value: strconv.Itoa(randomNum),Expires: time.Now().Add(72 * time.Hour),

		}
		http.SetCookie(w, cookie)

	}

	fmt.Println(cookie)

	//assign values 
	gues, _ := strconv.Atoi(r.FormValue("guess"))

	cookValue, _ := strconv.Atoi(cookie.Value)

	//create a new variable and assign the struct and string value
	 mess := &messagStruct{
		 Message: "Guess a number between 1 and 20", Guess: gues, Winner: false}

	//if the target is = to guess then declare the winner and change the target
	/*if cookValue == gues{

	 cookie = &http.Cookie{Name: "target", Value: strconv.Itoa(randomNum),Expires: time.Now().Add(72 * time.Hour),

		}
		http.SetCookie(w, cookie)

	}*/
	if cookValue == gues{

		mess.NewMessage = "You have Guessed the right number"
		mess.Winner = true

		//set a new randon number for the game
		cookie = &http.Cookie{Name: "target", Value: strconv.Itoa(getRandumNum(1, 20)),}
		http.SetCookie(w, cookie)
	
	}else if gues < cookValue{
	  	mess.NewMessage ="Try Again guess too low"
	}else {
		mess.NewMessage ="Try Again guess too high"
	}
	
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

//random number generator 
func getRandumNum(min, max int) int{
// Random int will be different each program execution.
	rand.Seed( time.Now().UnixNano())
   	return rand.Intn(max - min) + min
}

func main() {
	///////////open localhost on port 8080
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handleGuess)
	http.ListenAndServe(":8080", nil)
}