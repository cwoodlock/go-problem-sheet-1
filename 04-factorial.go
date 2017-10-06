package main

import "fmt"

//Main function
func main() {

	var num int
	fact := 1
	fmt.Println("Please enter a number: ")
	fmt.Scanf("%d", &num)

	if num == 0 {
		fmt.Printf("%d factorial is 1", num)
	} else {
			for i := num; i > 0; i-- {
				fact *= i
			}
	}



}
