package main  // this declaration shows that this is a standalone executable program

import (
	"fmt"
	"practiseGo/reverse"
)
func main() {
	res := reverse.Reverse("golang")
	fmt.Printf("The reverse is: %v", res);
	// if err != nil {
	// 	fmt.Printf(err)
	// } else {
	// 	fmt.Printf("The reverse is: %s", res);
	// }

}