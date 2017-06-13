package pointing

import (
  "fmt"
)
/*
Go vvariables are passed by value
*/
func Pointing() {
  module := 23
  ptr := &module
  lang := "Javascript"

  fmt.Println("\nThe memory address of module is", ptr)
  fmt.Println("\nThe value of module is", *ptr)
  changeLang(&lang)
  fmt.Println("\nWe started programming in", lang) // lang variable is unchanged
}

func changeLang(lang *string) string {
  *lang = "Go" // notice we don't use shorthand since we are not declaring a new variable
  fmt.Println("\nnWe are now programming in", *lang)
  return *lang
}
