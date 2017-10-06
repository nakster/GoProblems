//Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to 
//approximate the square root of a number x by picking a starting point z and then repeating the following operation.
package main

import (
    "fmt"
    "math"
)

func main() {
    //declare variables
    var i int

    //ask the user what number they would like to get the sqrt
    fmt.Print("Please enter a number to calculate Sqrt using Newton’s method: ")
    fmt.Scanf("%d",&i)
    
    //this prints out the values
     fmt.Println(i,"Squared in Sqrt Func is: ",Sqrt(i))
     fmt.Println(i,"Squared in Newton’s method is: ",Newt(i))
     fmt.Println("The Difference between the two is: ",math.Abs(Sqrt(i)-Newt(i)))
    
}

func Newt(newt int) float64 {

    //this converts the int value of user input to float64
    newNewt := float64(newt)

    if newNewt == 0 { return 0 }
    //floating point value 
    z := 1.0
    //for loop to get the sqrt
    for i := 0; i < int(newNewt); i++ {
        z = z - ((math.Pow(z, 2) - newNewt) / (2 * z))
    }
    //return sqrt
    return z
}

//function which uses math.sqrt to calculate the sqrt
func Sqrt(sqrt int) float64 {
    //this converts the int value of user input to float64
    newSqrt := float64(sqrt)
    //calculates the value
    return math.Sqrt(newSqrt)
}