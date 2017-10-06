//Largest and Smallest exercise in Go
//Author: Colm Woodlock
//Student Number: G00341460

//Watched parts of this in preparation: https://www.youtube.com/watch?v=bwId7l7aPeg

package main

import (
  "fmt"
)

func main() {

  //Creates a list
  list := []int{7, 10, 5, 8, 52, 1, 23}

  //Variables
  min := list[0]
  max := list[0]

  //For loop will loop for the length of the list
  for i := 0; i < len(list); i++ {
    //if statement will check if the current position of the list is less than min if it is assign it to min
    if list[i] < min {
      min = list[i]
    }
    //if statement will check if the current position of the list is larger than max if it is assign it to max
    if list[i] > max {
      max = list[i]
    }
  }
  //print result
  fmt.Printf("The min value of the list is: %d \nThe max value of the list is: %d", min, max)
}
