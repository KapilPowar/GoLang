package main

import (
	"fmt"
)

func main() {

	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultCh1 := make(chan int)
	resultCh2 := make(chan int)

	go splitter(resultCh1, resultCh2, data1, data2...)

	even := <-resultCh1
	odd := <-resultCh2
	finalResult := even + odd
	fmt.Println("Sum of even  == ", even)
	fmt.Println("Sum of  odd == ", odd)
	fmt.Println("Result == ", finalResult)

}

func splitter(resultch1 chan int, resultch2 chan int, nos1 []int, nos2 ...int) {
	odd, even := SumOf_OddEven(nos1)
	odd1, even1 := SumOf_OddEven(nos2)

	resultch1 <- even + even1
	resultch2 <- odd + odd1
}

func SumOf_OddEven(x []int) (odd int, even int) {
	for _, no := range x {
		if no%2 == 0 {
			even += no
		} else {
			odd += no
		}
	}
	return odd, even
}
