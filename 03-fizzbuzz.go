//FizzBuzz exercise in Go
//Author: Colm Woodlock

//Adapted from http://wiki.c2.com/?FizzBuzzTest

package main

import "fmt"

//Main function
func main() {
	i := 1
	for i < 50 {
		if i%3 == 0 && i%5 == 0 { //If statement to check if it is divisible by 3 and 5
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 { //If statement to check if it is divisible by only 5
			fmt.Println("Buzz")
		} else if i%3 == 0 { //If statement tot check if it is only divisible by 3
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
		i++ //increment i
	}
}
