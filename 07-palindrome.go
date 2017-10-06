//Palindrome exercise in Go
//Author: Colm Woodlock
//Student Number: G00341460

package main

import(
  "fmt"
)

func main(){

  //Declare variables
  var input string

  //Prompt user to enter string and read it from CLI
  fmt.Println("Please enter a string:")
  fmt.Scanf("%s \n", &input)

  //Declare the middle and last value of the string
  middle := len(input)/2
  last := len(input) -1

  //This checks to see if in the loop, the first nd last letters to see if they are not equal
  //Print if it is or isnt a palindrome
  for i := 0; i < middle; i++{
    if input[i] != input[last-i]{
      fmt.Printf("%s is not a Palindrome. \n", input)
    } else {
      fmt.Printf("%s is a Palindrome. \n", input)
    }
  }
}
