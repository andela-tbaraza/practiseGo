package structures

import (
    "fmt"
)

// normal function declaration
func my_func() int {
   //code
}

// func declaration with a type
type my_type struct { }

func (m my_type) my_func() int {
   //code
}
func Structures() {
    /* 
    structs - alllows us to define custom data types
    define types so that we can use later
    application of using small reusable components Go and OO
    Go has no objects, classes or inheritance
    Defining struct
    Structs can also contain behaviors in formof methods.
    The method definition is same as func only that you have to specify the type
    */



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