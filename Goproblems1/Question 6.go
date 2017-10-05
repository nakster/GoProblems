//Write a function that returns the largest and smallest elements in a list.
package main

import (
	"fmt"
	
)

func main() {
	//create a int array of 5 ints 
  var store [5]int
  //ask user for list of values
  fmt.Print("Please enter 5 Random Numbers\n")
  
  //populate array with values given
  for i:=0; i<5;i++{
	  //ask user for array values
	fmt.Print("Please enter number: ")
	fmt.Scanf("%d ",&store[i])
  }
  
  fmt.Print("the smallest and largest elements in a list are: ")
  fmt.Println(MinMax(store))
  
}

//function to produce min and max values
func MinMax(array [5]int) (int, int) {
	//setting min and max to first index of array
	var max int = array[0]
	var min int = array[0]
	//iterating over array to find min and max values
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
	}
    return min, max
}

