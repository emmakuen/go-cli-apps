package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	shouldCountLines := *flag.Bool("l", false, "Count lines")
	shouldCountBytes := *flag.Bool("p", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, shouldCountLines, shouldCountBytes))
}

func count(reader io.Reader, shouldCountLines bool, shouldCountBytes bool) int {
	scanner := bufio.NewScanner(reader)

	if !shouldCountLines && !shouldCountBytes {
		scanner.Split(bufio.ScanWords)
	} else if shouldCountBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
