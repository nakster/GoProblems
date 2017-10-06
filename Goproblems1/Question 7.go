//Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.
//http://www.golangpro.com/2016/01/check-string-palindrome-go.html
package main

import (   
    "fmt"
    "strings"
)

func main() {
    //declare variables
    var check string
    //ask the user for the string
    fmt.Print("Please Enter a string to check if its a palindrome:")
    fmt.Scanf("%s\n", &check)
    fmt.Println()
    //this checks if its palindrome or not 
     if palindrome(check) == true{
            fmt.Println("YES Its A Palindrome\n")
        } else {
            //otherwise return no
            fmt.Println("No Its Not A Palindrome\n")
        }  
}

//this function checks if the string is a palindrome
func palindrome(check string) bool{

    //this converts it to lowercase 
    check = strings.ToLower(check)

    //divide the string into half
    mid := len(check) / 2
    last := len(check) - 1
        
        //for 
        for i := 0; i < mid; i++ {
            //this checks if they are not Palimdrome
             if check[i] != check[last-i] {
                return false
             }
        }  
    //return true if it is
    return true
}
