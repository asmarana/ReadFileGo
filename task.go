package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CountWords(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("There is an error ! ", err)
	}
	defer file.Close()

	wordScans := bufio.NewScanner(file)
	wordScans.Split(bufio.ScanWords)
	var word, vowel int
	for wordScans.Scan() {
		word++
		var words string = strings.ToLower(wordScans.Text())

		if (words == ("a")) || (words == ("e")) || (words == ("i")) || (words == ("o")) || (words == ("u")) {
			vowel++
		}
	}
	return word, vowel

}

func CountLines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	lineScans := bufio.NewScanner(file)
	lineScans.Split(bufio.ScanLines)
	var Line int
	for lineScans.Scan() {
		Line++
	}
	return Line

}
func countChar(filename string) (int, int, int) {
	file, err := os.Open(filename)
	if err != nil {

		fmt.Println(err.Error())
	}
	charScans := bufio.NewScanner(file)
	charScans.Split(bufio.ScanBytes)
	var char, punc, integers int

	for charScans.Scan() {
		if charScans.Text() != " " {
			char++
		}
		if (charScans.Text() == (".")) || (charScans.Text() == (",")) || (charScans.Text() == (";")) || (charScans.Text() == ("!")) || (charScans.Text() == ("'")) {
			punc++
		}
		if regexp.MustCompile(`\d`).MatchString(charScans.Text()) {
			integers++
		}
	}
	return char, punc, integers

}

func main() {
	filename := "text.txt"
	char, punc, integers := countChar(filename)
	word, vowel := CountWords(filename)
	line := CountLines(filename)
	fmt.Printf("Numbers of Characters : %v\n", char)
	fmt.Printf("Numbers of Punctuations : %v\n", punc)
	fmt.Printf("Numbers of Integers : %v\n", integers)
	fmt.Printf("Numbers of Words : %v\n", word)
	fmt.Printf("Numbers of Lines : %v\n", line)
	fmt.Printf("Numbers of Vowels : %v\n", vowel)
}
