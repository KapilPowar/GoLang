package main

import "fmt"

func main() {
	counter := getCounter()
	fmt.Println(counter()) //=> 1
	fmt.Println(counter()) //=> 2
	fmt.Println(counter()) //=> 3
	fmt.Println(counter()) //=> 4
}

func getCounter() func() int { //=> step 1
	var count = 0           //=> step 2
	counter := func() int { //=> step 3
		count++ //=> step 4
		return count
	}
	return counter //=> step 5
}
