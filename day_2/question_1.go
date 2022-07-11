package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// ascii value of a is 97

func countCharacters(answer []int64, word string) {
	n := len(word)
	for i := 0; i < n; i++ {
		temp := word[i]
		atomic.AddInt64(&answer[temp-97], 1)
	}
}

func main() {
	words := make(chan string, 5)
	words <- "quick"
	words <- "brown"
	words <- "fox"
	words <- "lazy"
	words <- "dog"
	answer := make([]int64, 26, 26)

	go countCharacters(answer, <-words)
	go countCharacters(answer, <-words)
	go countCharacters(answer, <-words)
	go countCharacters(answer, <-words)
	go countCharacters(answer, <-words)

	time.Sleep(time.Second)

	j := 'a'
	for i := 0; i < 26; i++ {
		current := string(j)
		fmt.Printf("%v: %v ", current, answer[i])
		//fmt.Println()
		j++
	}
}
