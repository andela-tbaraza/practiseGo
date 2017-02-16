package main

import (
  "fmt"
)

func main () {
  // numbered lists conataining items of same type

  // Arrays vs slices
  // slices are references to the underlying array
  // slices make(type, len, cap)
  green := make([]int, 10, 20)
  green[0] = 1;
  fmt.Println(green) // [0 0 0 0 0 0 0 0 0 0]
  grass := []int{1, 2, 3, 4}
  fmt.Println(grass) //slice is a refernce to an array

  // append enables expansion of an array
  sluice := make([]int, 2, 4)
  fmt.Printf("The length of slice is %d and the capacity is %d", len(sluice), cap(sluice))
  mySlice := []int{3, 4, 5}

  // append doubles the capacity of array if exceeded
  for i := 0; i < 35; i++ {
    sluice = append(sluice, i)
    fmt.Printf("\nThe capacity is %d", cap(sluice))
  }

  // append to slices to slices using ellipses
  combine := append(sluice, mySlice...)
  fmt.Println("\nLord", combine)
}
