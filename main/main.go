package main

import (
	"fmt"
	"sync"
	fn "task"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	wg.Add(1)
	filename := "text.txt"
	c := make(chan int, 5)
	content := fn.ReadFile(filename)
	go fn.FileInfo(&c, content)
	fmt.Printf("Numbers of Words : %v\n", <-c)
	fmt.Printf("Numbers of Vowels : %v\n", <-c)
	fmt.Printf("Numbers of Integers : %v\n", <-c)
	fmt.Printf("Numbers of Lines : %v\n", <-c)
	fmt.Printf("Numbers of Punctuations : %v\n", <-c)

	wg.Wait()
	timetook := time.Since(start)
	fmt.Printf("Time in execution: %v\n", timetook)
}
