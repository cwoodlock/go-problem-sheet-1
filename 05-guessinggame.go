//Guessing game exercise in Go
//Author: Colm Woodlock
//Student Number: G00341460

//Adapted from: https://gobyexample.com/random-numbers

package main

import (
  "fmt"
  "math/rand"
)

func main(){

  //Declare variables
  tries := 0
  previous := 0
  guess := 0
  answer := rand.Intn(100)

  //Stay running until the guess is correct
  for guess != answer {
    fmt.Println("Please enter a number between 1 and 100: \n")
    fmt.Scanf("%d", &guess)

    //If guess is correct prompt user
    if guess == answer {
      tries++
      fmt.Printf("Correct! It took you %d times to guess the answer! \n", tries)
    } else if guess > answer {  //Else if guess is bigger propt user and count it as a try
      fmt.Printf("Your guess was too high please try again! \n")
      if previous != guess {  //If guess and the previous guess are the same do not count it as a try
        tries++
        previous = guess
      }
    } else if guess < answer {  //Else if guess is smaller propt user and count it as a try
      fmt.Printf("Your guess was too low please try again! \n")
      if previous != guess {  //If guess and the previous guess are the same do not count it as a try
        tries++
        previous = guess
      }
    }
  }

//ISSUE WITH CODE: rand.Intn(100) seems to only generate 81 and it prints lies 23 & 24 twice each loop

}
