package main

import "fmt"

//Main function
func main() {

	var num int
	fact := 1
	sum := 0
	fmt.Println("Please enter a number: ")
	fmt.Scanf("%d", &num)

	if num == 0 {
		fmt.Printf("%d factorial is 1", num)
	}
	for i := num; i > 0; i-- {
		fact *= i
	}

	fmt.Printf("%d factorial is: %d \n", num, fact)
	fmt.Printf("The sum of the constituant numbers of %d is: ",fact)
	
	for fact > 0{
		sum += fact % 10
		fact /= 10
	}

	fmt.Printf("%d",sum)


}
