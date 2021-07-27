package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Proident occaecat ex esse do duise mollit sunt id ea ut non consectetur incididunt Consectetur culpa non amet ipsum et ex duis Minim reprehenderit est officia est quis laborum laboris aute laboris laboris ut labore Lorem Reprehenderit culpa Lorem non duis nisi mollit cillum eiusmod aliquip est laborum mollit duis Aliqua eu tempor ad aliquip esse nostrud laborum excepteur ipsum excepteur eiusmod Sint nisi duis esse occaecat officia ut amet anim ex veniam est officia labore Elit esse exercitation officia quis ipsum adipisicing fugiat ex irure Culpa dolore laborum amet deserunt aliqua esse duis pariatur cillum est adipisicing Lorem eu Elit anim aute eiusmod consequat exercitation anim laborum consectetur excepteur adipisicing cillum Nostrud proident consequat cupidatat velit fugiat duis Ex tempor eu qui elit nostrud amet Nostrud dolore anim ex aliqua reprehenderit dolor velit adipisicing Consectetur eiusmod incididunt qui proident quis id reprehenderit Sit anim ea eu magna mollit laborum amet voluptate Dolore sunt laboris consectetur tempor nisi ut duis veniam cupidatat do Qui culpa ipsum incididunt sunt velit ipsum Est elit duis aliqua qui dolor aliqua enim laborum reprehenderit magna fugiat consectetur esse Quis nulla fugiat incididunt eu"

	//This is to get count of words occuring in sentence
	words := strings.Split(str, " ")
	wordCountBySize := getWordsCountBySize(words)
	wordSize, countLSize := getMaxWordsCountBySize(wordCountBySize)
	fmt.Printf("The size of the word that occurs the most(%d times) is %d\n", countLSize, wordSize)

	wordCount := getWordsCount(words)
	word, count := getMaxWordsCount(wordCount)
	fmt.Printf("The word that occurs the most(%d times) is %s \n", count, word)
}

func getWordsCountBySize(words []string) map[int]int {
	wordCountBySize := map[int]int{}
	for _, word := range words {
		wordCountBySize[len(word)] += 1
	}
	return wordCountBySize
}

func getWordsCount(words []string) map[string]int {
	wordCount := map[string]int{}
	for _, word := range words {
		wordCount[word] += 1
	}

	return wordCount
}

func getMaxWordsCountBySize(wordCountBySize map[int]int) (wordSize, count int) {
	for size, cnt := range wordCountBySize {
		if count < cnt {
			count = cnt
			wordSize = size
		}
	}
	return
}

func getMaxWordsCount(wordCount map[string]int) (word string, count int) {
	maxcount, maxword := 0, ""
	for i, w := range wordCount {
		//fmt.Println(i, w)
		if maxcount < w {
			maxcount = w
			maxword = i
		}
	}
	return maxword, maxcount
}
