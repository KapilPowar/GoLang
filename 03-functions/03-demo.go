package main

import "fmt"

func main() {

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(100, 200))

	//anonymous func

	func() {
		fmt.Println("Printed form anonymous ")
	}()

	func(x, y int) {
		fmt.Println("Printed form anonymous ", x-y)
	}(100, 200)

	subtractResult := func(x, y int) int {
		return x - y
	}(100, 200)
	fmt.Println("Printed form anonymous substract ", subtractResult)
}
