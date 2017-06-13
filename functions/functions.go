package functions

import (
	"fmt"
	"strings"
)

func Functions() {
	book := "I kissed dating goodbye" //declare and initialize variables
	author := 32
	one := 4
	two := 2
	fmt.Println(one%two == 0)

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
