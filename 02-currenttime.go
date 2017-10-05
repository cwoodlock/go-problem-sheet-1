//Print the current time in Go
//Author: Colm Woodlock

//Adapted from https://tour.golang.org/welcome/4

package main

import (
	"fmt"
	"time"
)

//Main function
func main() {
	//Prints out the current time
	fmt.Println("The time is", time.Now())
}
