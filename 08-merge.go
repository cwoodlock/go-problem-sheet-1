//Merge exercise in Go
//Author: Colm Woodlock

//Adapted from https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go


package main

import(
  "fmt"
  "sort"
)

func main(){

  //Declare lists to merge
  list := []int{1,4,6}
  list2 := []int{2,3,5}

  //List that will contain the merged list
  list3 := []int{1}

  //Apend the lists together
  list3 = append(list,list2...)

  //Sort the list3
  sort.Ints(list3)

  //Print lists and merged and sorted list
  fmt.Println("List 1:",list)
  fmt.Println("List 1:",list2)
  fmt.Println("The merged list is: ",list3)
}
