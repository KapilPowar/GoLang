package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go add(10, 20, ch)
	go add(10, 10, ch)
	result := <-ch
	fmt.Println("result = ", result)
	result = <-ch
	fmt.Println("result = ", result)
	wg.Wait()
	fmt.Println("done")
}
func add(x, y int, ch chan int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
