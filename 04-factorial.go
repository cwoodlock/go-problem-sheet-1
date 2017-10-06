//Factorial exercise in Go
//Author: Colm Woodlock
//Student Number: G00341460

package main

import "fmt"

//Main function
func main() {

	//Decalre variables
	var num int
	fact := 1
	sum := 0

	//Prompt user to enter a number and read it in from the CLI
	fmt.Println("Please enter a number: ")
	fmt.Scanf("%d", &num)

	//Calculate the factorial of the number they entered check if it is 0
	if num == 0 {
		fmt.Printf("%d factorial is 1", num)
	}
	for i := num; i > 0; i-- {
		fact *= i
	}

	//Print results
	fmt.Printf("%d factorial is: %d \n", num, fact)
	fmt.Printf("The sum of the constituant numbers of %d is: ",fact)

	//The in this loop you get the modulus 10 of the factorial and add it to sum,
	//Then you divide factorial by 10 and the loop repeats, this gets the sum of the constituant numbers
	for fact > 0{
		sum += fact % 10
		fact /= 10
	}

	//Print out the sum of the constituant numbers of the factorial
	fmt.Printf("%d",sum)


}
