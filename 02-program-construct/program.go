package main

import "fmt"

func main() {
	x := 23
	if x%2 == 0 {
		fmt.Println(x, " is even number")
	} else {
		fmt.Println(x, " is 0dd number")
	}

	if y := 33; y%2 == 0 {
		fmt.Println(y, " is even number")
	} else {
		fmt.Println(y, " is 0dd number")
	}

}
