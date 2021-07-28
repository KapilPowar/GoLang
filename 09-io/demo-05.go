package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	filecontents, err := ioutil.ReadFile("data2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(filecontents))
}
