package main

import (
	"fmt"
	"io/ioutil"
	// "unicode"
)

func readFile(filename string) string {
	context, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There is an error", err)
	}
	return string(context)
}

func countWords(context string) int {

	var words int
	for _, i := range context {
		if is_word(i) {
			words++
		}
	}
	return words
}

func countVowel(context string) int {
	var vowel int
	for _, i := range context {
		switch i {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel++
		}
	}
	return vowel
}

func countIntegers(context string) int {
	var integers int
	for _, i := range context {
		// unicode.IsDigit(i)
		switch i {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			integers++
		}
	}
	return integers
}

func countLines(context string) int {
	var lines int
	for _, i := range context {
		if is_line(i) {
			lines++
		}
	}
	return lines
}

func countPunc(context string) int {
	var punc int
	for _, i := range context {
		switch i {
		case ',', '.', ';', '"', '=', '!', '/':
			punc++
		}
	}
	return punc
}

func is_word(i rune) bool {
	if i == ' ' {
		return true
	} else {
		return false
	}
}

func is_line(i rune) bool {
	if i == '\n' {
		return true
	} else {
		return false
	}
}

func main() {
	filename := "text.txt"
	context := readFile(filename)
	words := countWords(context)
	vowels := countVowel(context)
	integers := countIntegers(context)
	lines := countLines(context)
	punc := countPunc(context)

	fmt.Printf("Numbers of Words : %v\n", words)
	fmt.Printf("Numbers of Vowels : %v\n", vowels)
	fmt.Printf("Numbers of Integers : %v\n", integers)
	fmt.Printf("Numbers of Punctuations : %v\n", punc)
	fmt.Printf("Numbers of Lines : %v\n", lines)

}
