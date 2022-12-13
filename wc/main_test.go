package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	wordBuffer := bytes.NewBufferString("w1 w2 w3 w4\n")
	expected := 4
	result := count(wordBuffer, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	wordBuffer := bytes.NewBufferString("w1 w2 w3 w4\nline2\nline3 w1")

	expected := 3
	result := count(wordBuffer, true)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}
