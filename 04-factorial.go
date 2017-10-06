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
	for fact > 0{
		sum += fact % 10
		fact /= 10
	}


}
