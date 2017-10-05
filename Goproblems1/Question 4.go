// //Program to find Factorial of number
// //Write a function that calculates the sum of the digits of the factorial of a number. n! means n x (n − 1) … x 3 x 2 x 1. 
// //For example, 10! = 10 x 9 x … x 3 x 2 x 1 = 3628800, and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. 
// //Find the sum of the digits in the number 100! 
package main
import (
	"fmt"
)

//this main method asks for the input 
func main(){  
    // Declare Variable
    var num int

    //asks the user for the input  
    fmt.Print("Please Enter A Number between 0 and 100 : ")
    fmt.Scan(&num)
    
    //this line uses the methods to calculate the factorial 
	//pass the int to the function
    if num > 0 && num < 21{
        fmt.Print("Factorial is: ",CalFacOfInt(num))
        fmt.Println();
    }else{
         fmt.Print("Factorial is: ",Calcfactorial(num))
         fmt.Println();
    }
    
    fmt.Print("The sum of factorial is: ",P20(num))
}//eomain

//function to calculate 
func P20(n int) int {
    //declare sum
    sum := 0;
    //this stores the int values
    arrayForInt := [200]int{};
    //intialise it
    arrayForInt[0] = 1;
    //for loop
    for i := 2; i <= n; i++ {
    	for j := 0; j < len(arrayForInt); j++ {
    		arrayForInt[j] *= i;
    		if j > 0 && arrayForInt[j - 1] > 9 {
    			arrayForInt[j] += int(arrayForInt[j - 1] / 10);
    			arrayForInt[j - 1] %= 10;
    		}
    	}
    }
    for i := 0; i < len(arrayForInt); i++ {
    	sum += arrayForInt[i];
    }
    //return the value
    return sum;
}

//var is declaring a variable 
//CalculateNum is the name of the variable and float64 is the type 
// uint64 ranges between 0 through 18446744073709551615
// Use float64 when using big number or big calculations
//variables
var CalculateNum float64 = 1                       
var i int

//calculate factorial doesnt matter what you name the int inside the = (num int)
func Calcfactorial(num int) float64 {   
    if(num < 0){
        fmt.Print("Can't calculate Minus Numbers.")    
    }else{        
        for i:=1; i<=num; i++ {
			// mismatched types float64 and int
            CalculateNum *= float64(i)  
        }//eof
    }  
	// return from method
    return CalculateNum  
}
//variables
var CalculateNumInt uint64 = 1                       
var j int

//this function calulates the factorial upto 20
func CalFacOfInt(num int) uint64 {   
    if(num < 0){
        fmt.Print("Can't calculate Minus Numbers.")    
    }else{        
        for j:=1; j<=num; j++ {
			// mismatched types float64 and int
            CalculateNumInt *= uint64(j)  
        }//eof
    }  
	// return from method
    return CalculateNumInt  
}