package task

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There is an error", err)
	}
	return string(content)
}

func FileInfo(c *chan int, content string) {

	var words, vowel, integers, lines, punc int
	for _, char := range content {
		if isWord(char) {
			words++
		}
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel++
		}
		switch char {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			integers++
		}
		if isLine(char) {
			lines++
		}
		switch char {
		case ',', '.', ';', '"', '=', '!', '/', '~', '@', '#', '%', '&', '?', ':':
			punc++
		}
	}
	*c <- words
	*c <- vowel
	*c <- integers
	*c <- lines
	*c <- punc
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
