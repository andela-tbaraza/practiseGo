package main

import (
    "fmt"
)

func main() {
    // structs - alllows us to define custom data types
    // define types so that we can use later
    // application of using small reusable components Go and OO
    // Go has no objects, classes or inheritance
    // Defining struct

     type circle struct {
        Radius float64
        Area float32
        Perimeter float64
    }

    // var example2 circle
    // example2 := new(circle)
    // fmt.Println(example2)


    example := circle{
        Radius: 34.555,
        Area: 43333.444,
        Perimeter: 3432333,
        
    }

    fmt.Println(example)
    fmt.Println(example.Radius)


}