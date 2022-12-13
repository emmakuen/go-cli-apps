package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("p", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *countLines, *countBytes))
}

func count(reader io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(reader)

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
