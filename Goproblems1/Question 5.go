//Write a guessing game where the user must guess a secret number. After every guess the program tells the user 
//whether their number was too large or too small. At the end the number of tries needed should be printed.
//It counts only as one try if they input the same number multiple times consecutively.
//http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html
package main

import (
	"fmt"
	"math/rand"
    "time"  
)

func main() {

    // Declare variables 
    count := 0
    var AskUser int
	randomNum := getRandumNum(1, 20)
    var guessedNum [] int

	//this for loop prints outs the results 
    for count < 20 {
        //ask user what number it may be
        fmt.Print("Enter a Number Between 1 and 20 To Guess: ")
        fmt.Scanf("%d ", &AskUser)
		//this keeps the amount of time the user guessed 
        count++
        //check the guess and see if its higher or lower
        if AskUser < randomNum {
            fmt.Println("Its Too Low")
        } else if AskUser > randomNum {
            fmt.Println("Its Too High")
        } else {
            break
        }
        //check if it already exits in array
        if contains(guessedNum, AskUser) == true{
            fmt.Printf("Already Entered Before\n")
            count--;
        } else {
            //otherwise add to array/slice
            guessedNum = append(guessedNum, AskUser)      
        }  
    }
    if AskUser == randomNum {
        fmt.Printf("You Guessed in %d tries\n", count)
    } else {
        fmt.Printf("Nope. The number was %d\n", randomNum)
    }
}

//random number generator 
func getRandumNum(min, max int) int{
// Random int will be different each program execution.
	rand.Seed( time.Now().UnixNano())
   	return rand.Intn(max - min) + min
}

//checks for duplicates //return type boolean 
func contains(s []int, e int) bool {
    //iterate to find the value if contains or else
    for _, a := range s {
        if a == e {
            //if returns true its a duplicate
            return true
       }
    }
    //if its not true then it wasnt asked 
    return false
}