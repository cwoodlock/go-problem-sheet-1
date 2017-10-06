//Adapted from: https://gobyexample.com/random-numbers

package main

import (
  "fmt"
  "math/rand"
)

func main(){

  tries := 0
  previous := 0
  guess := 0
  answer := rand.Intn(100)

  for guess != answer {
    fmt.Println("Please enter a number between 1 and 100: \n")
    fmt.Scanf("%d", &guess)

    if guess == answer {
      tries++
      fmt.Printf("Correct! It took you %d times to guess the answer! \n", tries)
    } else if guess > answer {
      fmt.Printf("Your guess was too high please try again! \n")
      if previous != guess {
        tries++
        previous = guess
      }
    } else if guess < answer {
      fmt.Printf("Your guess was too low please try again! \n")
      if previous != guess {
        tries++
        previous = guess
      }
    }
  }


}
