package main

import (
    "fmt"
    "regexp"
)

/*
given a string containing only parentheses (i.e. "(()(" or "(())")
return whether or not the all parentheses are matched

(()) => true
))(( => false
() => true
((()())()) => true
*/

// check if each opening parentheses ( has a closing )

func parensMatch(s string) (parensDoMatch bool) {
    return
}

func main() {
    str := "()(())"
    // res := parensMatch(str)
    // fmt.Fprintln(res)
    matched, err := regexp.MatchString("^(\\(\\))+$", str)
    fmt.Println(matched, err)


}


