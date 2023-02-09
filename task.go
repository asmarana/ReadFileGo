package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("There is an error !", err)
	}

	defer file.Close()

	wordScans := bufio.NewScanner(file)
	wordScans.Split(bufio.ScanWords)
	var counter, vowel int
	for wordScans.Scan() {
		counter++
		if (wordScans.Text() == ("a")) || (wordScans.Text() == ("e")) || (wordScans.Text() == ("i")) || (wordScans.Text() == ("o")) || (wordScans.Text() == ("u")) {
			vowel++
		}
	}

	file1, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	lineScans := bufio.NewScanner(file1)
	lineScans.Split(bufio.ScanLines)
	var Line int
	for lineScans.Scan() {
		Line++
	}

	file2, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	charScans := bufio.NewScanner(file2)
	charScans.Split(bufio.ScanBytes)
	var char, punc int
	for charScans.Scan() {
		if charScans.Text() != " " {
			char++
		}
		if (charScans.Text() == (".")) || (charScans.Text() == (",")) || (charScans.Text() == (";")) || (charScans.Text() == ("!")) || (charScans.Text() == ("'")) {
			punc++
		}
	}

	file3, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	intScans := bufio.NewScanner(file3)
	intScans.Split(bufio.ScanBytes)
	var integers int
	for intScans.Scan() {
		if regexp.MustCompile(`\d`).MatchString(intScans.Text()) {
			integers++
		}
	}

	fmt.Printf("Numbers of Words : %v\n", counter)
	fmt.Printf("Numbers of Lines : %v\n", Line)
	fmt.Printf("Numbers of Characters : %v\n", char)
	fmt.Printf("Numbers of Vowels : %v\n", vowel)
	fmt.Printf("Numbers of Punctuations : %v\n", punc)
	fmt.Printf("Numbers of Integers : %v\n", integers)
}

func main() {
	filename := "text.txt"
	ReadFile(filename)
}
