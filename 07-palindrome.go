//Palindrome exercise in Go
//Author: Colm Woodlock

package main

import(
  "fmt"
)

func main(){

  var input string

  fmt.Println("Please enter a string:")
  fmt.Scanf("%s \n", &input)

  middle := len(input)/2
  last := len(input) -1

  for i := 0; i < middle; i++{
    if input[i] != input[last-i]{
      fmt.Printf("%s is not a Palindrome. \n", input)
    } else {
      fmt.Printf("%s is a Palindrome. \n", input)
    }
  }
}
