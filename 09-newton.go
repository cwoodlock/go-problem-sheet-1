//Newton exercise in Go
//Author: Colm Woodlock

//Adapted from https://tour.golang.org/flowcontrol/8 and code from lab

package main

import (
    "fmt"
    "math"
)

func z_next(z float64, x float64) float64
{
  return z - (((z * z) - x) / (2 * z))
}
func main() {

  //Variable x
  var x float64

  //Prompts user to enter a number and reads it from the CLI
  fmt.Println("Please enter a number to get the square root of: \n")
  fmt.Scanf("%f", &x)

  //the initial guess
  z := float64(1)

  //itterate until the next guess is the same as the current
  for z = 1.0; z != z_next(z,x); z = z_next(z, x){
    //Print the guess for each itteration
    fmt.Printf("Current guess: %1.28f\n", z)
  }

  //Finally, z is a good approzimation of the square root
  fmt.Printf("The square root of %f is %f", x, z)

  //Print out the math.Sqrt value
  fmt.Printf("math.Sqrt gives the value as %f: \n", math.Sqrt(x))

}
