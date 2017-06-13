package looping

import (
	"fmt"
	"time"
)

func Looping() {
	// we want to create a function that will count down from 10
	// we want to use time package
	// i := 10; - get executed before first evaluation of condition
	// 1-- gets executed after each execution of the block

	for i := 10; i >= 0; i-- {
		if i == 0 {
			fmt.Println("Boooooom")
			break // gets out of the current loop to the previous loop
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second) // delay for 1 second
	}

	// for range loop
	unorderedList := []string{"Mom", "dad", "siz", "bro"}
	fmt.Println(unorderedList)

	for _, i := range unorderedList {
		fmt.Println(i)
	}

	// break and continue
	for i := 10; i >= 0; i-- {
		if i%2 == 0 {
			fmt.Println("Boooooom")
			continue // alters the loop and goes to the top of the loop
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second) // delay for 1 second
	}

	// to get out of more than loop using break implement labels for break

	/*
		  for <expression> {
		    <code>
		  breakloop
		    for <expression> {
		      <code>
		      for <expression> {
		        <code>
		        break breakloop
		    }
		  }
		}
	*/
}
