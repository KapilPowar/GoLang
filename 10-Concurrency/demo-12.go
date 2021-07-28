package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		split the data into two sets, add them concurrently,
		collate the final result and display the final result
	*/
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}

	wg := &sync.WaitGroup{}
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	data1 := data[:len(data)/2]
	data2 := data[len(data)/2:]

	resultCh1 := make(chan int)
	resultCh2 := make(chan int)
	go sum(resultCh1, data1...)
	go sum(resultCh2, data2...)
	first := <-resultCh1
	second := <-resultCh2
	finalResult := first + second
	fmt.Println("Result of first half 1== ", first)
	fmt.Println("Result of 2nd half 1== ", second)
	fmt.Println("Result 1== ", finalResult)

	//another solution
	wg.Add(2)
	go add(data1, 2, ch1, ch2, wg)
	go add(data2, 2, ch2, ch1, wg)
	ch1 <- 1
	ch2 <- 1
	wg.Wait()

	fmt.Println("Result2 == ", <-ch1+<-ch2)
	// fmt.Println("Done")
}

func add(d []int, delay time.Duration, in chan int, out chan int, wg *sync.WaitGroup) {
	result := 0
	for i := 0; i < len(d); i++ {
		<-in
		//println(d[i])
		result += d[i]
		out <- result
	}
	wg.Done()
}

func sum(resultCh chan int, nos ...int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	resultCh <- result
}
