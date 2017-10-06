Q7 
package main

import (   
    "fmt"
    "strings"
)

func main() {

    var palindrome string
    fmt.Print("Please Enter a string to check if its a palindrome:")
    fmt.Scanf("%s\n", &palindrome)
     palindrome = strings.ToLower(palindrome)
     fmt.Println(isP(palindrome))
}
//Function to test if the string entered is a Palindrome

func isP(s string) string {

    //s = 

    mid := len(s) / 2
    last := len(s) - 1
        for i := 0; i < mid; i++ {
             if s[i] != s[last-i] {
                return "NO. It's not a Palimdrome."
             }
        }
    
    return "YES! You've entered a Palindrome"
}