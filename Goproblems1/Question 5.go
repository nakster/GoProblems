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
	myrand := xrand(1, 20)
    var guessedNum [] int

	//this for loop prints outs the results 
    for count < 20 {
        //ask user what number it may be
        fmt.Println("Guess a number between 1 and 100.: ")
        fmt.Scanf("%d ", &AskUser)
		//this keeps 
        count++
        //check if guess is higher or lower than random number
        if AskUser < myrand {
            fmt.Println("Your guess is too low.")
        } else if AskUser > myrand {
            fmt.Println("Your guess is too high.")
        } else {
            break
        }
        //check if guess is in slice
        if contains(guessedNum, AskUser) == true{
            fmt.Printf("That number has already been entered\n")
            count--;
        } else {
            //add guess to slice of guessed numbers
            guessedNum = append(guessedNum, AskUser)
            
        }  
    }
    if AskUser == myrand {
        fmt.Printf("Good job! You guessed the correct number in %d tries\n", count)
    } else {
        fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
    }
}

//this generates random number between given range
func getRandumNum(min, max int) int{
// Random int will be different each program execution.
// rand.Seed(time.Now().Unix())
	rand.Seed( time.Now().UnixNano())
   	return rand.Intn(max - min) + min

}

//this generates random number between given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
//this checks slice of guessed numbers and doesnt count duplicates
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true //integer is in list
        }
    }
    return false
}