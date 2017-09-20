//Write a guessing game where the user must guess a secret number. After every guess the program tells the user 
//whether their number was too large or too small. At the end the number of tries needed should be printed.
//It counts only as one try if they input the same number multiple times consecutively.
package main

import (
	"fmt"
	"math/rand"
    "time"
  
)

func main() {

	var guess int

	//println will ask for the number on the next line use print
	fmt.Print("Please Enter the Number to Play the Game : ")
	fmt.Scan(&guess)
	

}//main

func getRandumNum(){

// Random int will be different each program execution.
	rand.Seed( time.Now().UnixNano())
    var bytes int

    bytes = rand.Intn(100)+1
    fmt.Println(bytes)

	fmt.Print("Random Number is ",bytes)
	
	checkNum(bytes)

}

func checkNum(a int){

	return a
	
}

