package main // this declaration shows that this is a standalone executable program

import (
  "fmt"
  "strings"
)

func main () {
  book := "I kissed dating goodbye" //declare and initialize variables
  author := 32

  fmt.Println(titlecase(book, author)) // call the titlecase function and pass in copy of
                                      // book and author
}

func titlecase(book string, author int) (s1 string, s2 int) {
  // function signature
  book = strings.Title(book)
  author = 34

  return book, author
}

// variadic functions
