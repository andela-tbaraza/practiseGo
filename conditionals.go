package main

import (
  "fmt"
  "os"
)

func main () {
  // lady := 23
  // man := 30

  // first match in lexical order gets executed ang the if blocked is jumped

  if lady, man := 17, 30; lady < man {
    // using optional inititialization
    // their scope is on the if block only
    fmt.Println("There is potential for a meaningful relationship")
    // we can also have nested if
    if lady < 20 {
      fmt.Println("That's an early marriage, offense")
    }
  } else if lady > man {
    fmt.Println("The lady has the potential of being a sugar mummy")
  } else {
    fmt.Println("I have no idea what's going on")
  }

  // SWITCH
  switch lady, man := 25, 30; lady < man {
  case true:
    fmt.Println("There is potential for a meaningful relationship")
  case false:
    fmt.Println("The lady has the potential of being a sugar mummy")
  default:
    fmt.Println("I have no idea what's going on")
  }

  // Break and fall-through in SWITCH

  switch "grass" {
  case "car":
    fmt.Println("This is a car")
  case "gun", "grass":
    fmt.Println("This is a gun")
  case "house":
    fmt.Println("This is a house")
    fallthrough // apply for a case by case better use multiple values in the case as line 43
  case "money":
    fmt.Println("This is money")
  default:
    fmt.Println("Input an item")
  }
  // if Error handling

  _, err := os.Open("Tonie/my-projects")
  if err != nil {
    fmt.Println("The error is", err)
  }

}
