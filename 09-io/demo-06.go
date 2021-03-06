package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()

	bufferedReader := bufio.NewReader(fileReader)
	sentenceCount := 0
	for {
		sentence, err := bufferedReader.ReadString('.')
		if err == io.EOF {
			fmt.Println("EOF thats all to read!!!")
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		sentenceCount++
		fmt.Printf("sentence # %d: %s\n", sentenceCount, sentence)
	}
}
