//Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to 
//approximate the square root of a number x by picking a starting point z and then repeating the following operation.
package main

import (
    "fmt"
    "math"
)

func main() {

    var i int

    fmt.Print("Please enter a number to calculate Sqrt using Newton’s method: ")
    fmt.Scanf("%d",&i)
 
     sqrt := Sqrt(float64(i))
     newt := Newt(float64(i))
     fmt.Println(i, "squared:")
     fmt.Println("  Sqrt:", sqrt)
     fmt.Println("  Newt:", newt)
     fmt.Println("  Difference:", math.Abs(sqrt-newt))
    
}

func Newt(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}