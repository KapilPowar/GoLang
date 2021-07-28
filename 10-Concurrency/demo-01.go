package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Hello")
	go print("world")
	fmt.Println("done")
	time.Sleep(2 * time.Second)
	//runtime.Gosched()
}

func print(s string) {
	fmt.Println(s)
}
