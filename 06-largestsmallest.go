//Largest and Smallest exercise in Go
//Author: Colm Woodlock

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

  for i := 0; i < len(list); i++ {
    if list[i] < min {
      min = list[i]
    }
    if list[i] > max {
      max = list[i]
    }
  }

  fmt.Printf("The min value of the list is: %d \nThe max value of the list is: %d", min, max)
}
