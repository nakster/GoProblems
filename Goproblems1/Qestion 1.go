import (
	"fmt"
	"time"
)

func main() {

	//(1Write a program that prints “Hello, world!” in Japanese (using Japanese 		
	//characters) to the screen.
	fmt.Println("こんにちは, 世界")
	println("It says Hello World")
	
	//(2 Write a program that prints the current time and date to the console.
	fmt.Println("The time is", time.Now())
	
	//(3Write a program that prints the numbers from 1 to 100, except for the 	
	//following conditions. For multiples of three print “Fizz” instead of the 	
	//number, and for the multiples of five print “Buzz”. For numbers which are	
	//multiples of both three and five print “FizzBuzz”.
	
	for i:=1; i<=100; i++ {//cant have spaces between if statements//sof
		if i % 3 == 0 && i % 5 == 0{
			fmt.Println("FizzBuzz")
		}else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}//eof
	
	
}//main