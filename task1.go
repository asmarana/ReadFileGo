package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var wg sync.WaitGroup

func readFile(filename string) string {
	context, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There is an error", err)
	}
	return string(context)
}

func countWords(c *chan int, context string) {

	var words int
	for _, char := range context {
		if isWord(char) {
			words++
		}
	}
	*c <- words
	defer wg.Done()
}

func countVowel(c1 *chan int, context string) {
	var vowel int
	for _, char := range context {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel++
		}
	}
	*c1 <- vowel
	defer wg.Done()
}

func countIntegers(c2 *chan int, context string) {
	var integers int
	for _, char := range context {
		switch char {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			integers++
		}
	}
	*c2 <- integers
	defer wg.Done()
}

func countLines(c3 *chan int, context string) {
	var lines int
	for _, char := range context {
		if isLine(char) {
			lines++
		}
	}
	*c3 <- lines
	defer wg.Done()
}

func countPunc(c4 *chan int, context string) {
	var punc int
	for _, char := range context {
		switch char {
		case ',', '.', ';', '"', '=', '!', '/', '~', '@', '#', '%', '&', '?', ':':
			punc++
		}
	}
	*c4 <- punc
	defer wg.Done()
}

func isWord(char rune) bool {
	if char == ' ' {
		return true
	} else {
		return false
	}
}

func isLine(char rune) bool {
	if char == '\n' {
		return true
	} else {
		return false
	}
}

func main() {
	start := time.Now()
	wg.Add(5)
	filename := "text.txt"
	c := make(chan int, 1)
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	c4 := make(chan int, 1)
	context := readFile(filename)
	go countWords(&c, context)
	go countVowel(&c1, context)
	go countIntegers(&c2, context)
	go countLines(&c3, context)
	go countPunc(&c4, context)
	fmt.Printf("Numbers of Words : %v\n", <-c)
	fmt.Printf("Numbers of Vowels : %v\n", <-c1)
	fmt.Printf("Numbers of Integers : %v\n", <-c2)
	fmt.Printf("Numbers of Lines : %v\n", <-c3)
	fmt.Printf("Numbers of Punctuations : %v\n", <-c4)

	wg.Wait()
	timetook := time.Since(start)
	fmt.Printf("Time in execution: %v\n", timetook)
}
