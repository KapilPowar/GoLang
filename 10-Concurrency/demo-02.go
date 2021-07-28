package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(2)
	go print("Hello")
	go print("world")
	fmt.Println("done")
	wg.Wait()

}

func print(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
	wg.Done()
}
