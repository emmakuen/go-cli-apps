package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	wordBuffer := bytes.NewBufferString("w1 w2 w3 w4\n")
	expected := 4
	result := count(wordBuffer)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}
