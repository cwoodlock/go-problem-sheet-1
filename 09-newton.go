//Newton exercise in Go
//Author: Colm Woodlock

//Adapted from https://tour.golang.org/flowcontrol/8 and code from lab

package main

import (
    "fmt"
    //"math"
)

func main() {

  var x int

  fmt.Println("Please enter a number to get the square root of: \n")
  fmt.Scanf("%d", &x)

  fmt.Println(x)
}
