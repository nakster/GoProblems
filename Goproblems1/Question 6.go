//Write a function that returns the largest and smallest elements in a list.
package main

import (
	"fmt"	
)

func main() {
  //create a int array of 5 ints 
  var store [5]int
  //ask user for values
  fmt.Print("Please enter 5 Random Numbers\n")
  
  //put the values into the array/slice
  for i:=0; i<5;i++{
	  //ask user for array values
	fmt.Print("Please enter number: ")
	fmt.Scanf("%d ",&store[i])
  }
  //print out the results
  fmt.Println("The largest elements in a list is: ",Max(store))
  fmt.Println("The Smallest elements in a list is: ",min(store))
  
}

//gets the min value
func min(array [5]int) int{

    //set to 0
    var min int = array[0]

    //iterates to find min value
    for _, value := range array {
        if min > value {
            min = value
        }
	}
    return min
}
//gets the max value
func Max(array [5]int) int {
	//set to 0
	var max int = array[0]

	 //iterates to find max value
    for _, value := range array {
        if max < value {
            max = value
        }
	}
    return max
}
