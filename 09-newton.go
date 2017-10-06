//Newton exercise in Go
//Author: Colm Woodlock

//Adapted from https://tour.golang.org/flowcontrol/8 and code from lab

package main

import (
    "fmt"
    "math"
)

func main() {

  //Variable x
  var x float64

  //Prompts user to enter a number and reads it from the CLI
  fmt.Println("Please enter a number to get the square root of: \n")
  fmt.Scanf("%f", &x)

  //Prints the square root of user input
  fmt.Printf("square root of your entered value equals: %f",math.Sqrt(x))

}
