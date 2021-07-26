package main

import "fmt"

func main() {
	x := 10
	y := 20
	//z := 30
	// addResult := add(x, y)
	// fmt.Printf("Result of proccessing %d & %d is %d\n", x, y, addResult)
	// subtractResult := sub(x, y)
	// fmt.Printf("Result of proccessing %d & %d is %d\n", x, y, subtractResult)

	compute(x, y, add)
	compute(x, y, sub)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func compute(x, y int, fn func(int, int) int) {
	result := fn(x, y)
	fmt.Printf("Result of proccessing %d & %d is %d\n", x, y, result)

}
