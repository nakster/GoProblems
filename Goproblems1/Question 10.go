//Write a function to reverse a string in Go.
package main

import (   
    "fmt"
    
)

func main() {

    //declare variables
    var reverse string

    //ask the user for the string
    fmt.Print("Please Enter a string to Reverse:")
    fmt.Scanf("%s\n", &reverse)
    fmt.Println()
    //print out the reversed string
    fmt.Println("The String",reverse,"is reversed into:",Reverse(reverse))
    fmt.Println()

}

func Reverse(value string) string {
    
    // convert using runes
    data := []rune(value)
    result := []rune{}

    // add them in reverse
    for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }
    // Return new string
    return string(result)
}
