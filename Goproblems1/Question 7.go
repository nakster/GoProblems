//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted from http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html
//slice function adapted from https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
//Guessing game 

//import package
package main
import(
    "fmt"
    "math/rand"
    "time"
)

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

func main() {
    // Variable Declaration
    myrand := xrand(1, 100)
    guessTaken := 0
    var guess int
    var guessedNum [] int


    //this is the while loop
    for guessTaken < 100 {
        //ask user for guess
        fmt.Println("Guess a number between 1 and 100.: ")
        fmt.Scanf("%d ", &guess)
        guessTaken++
        //check if guess is higher or lower than random number
        if guess < myrand {
            fmt.Println("Your guess is too low.")
        } else if guess > myrand {
            fmt.Println("Your guess is too high.")
        } else {
            break
        }
        //check if guess is in slice
        if contains(guessedNum, guess) == true{
            fmt.Printf("That number has already been entered\n")
            guessTaken--;
        } else {
            //add guess to slice of guessed numbers
            guessedNum = append(guessedNum, guess)
            
        }
       
    }
    if guess == myrand {
        fmt.Printf("Good job! You guessed the correct number in %d tries\n", guessTaken)
    } else {
        fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
    }
}

