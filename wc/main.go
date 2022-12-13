package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(reader io.Reader, countByLine bool) int {
	scanner := bufio.NewScanner(reader)

	if !countByLine {
		scanner.Split(bufio.ScanWords)
	}

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
