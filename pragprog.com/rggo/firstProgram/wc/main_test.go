package main

import (
	"bytes"
	"testing"
      )

//TestCountWords tests the count function set to count words.
func TestCountWords(t *testing.T) { 
  // Create a new buffer of bytes from a string containing 4 words. 
  // Pass it to the count function. An 
  b := bytes.NewBufferString("word1 word2 word3 word4\n")

  exp := 4

  res := count(b, false)

  if res != exp {
	  t.Errorf("Expected %d, got %d instead.\n", exp, res)
	  }

}

//TestCountLines tests the number of lines
func TestCountLines(t *testing.T) {
  b := bytes.NewBufferString("word1 word2 word3\n line2\nline 3 word1")

  exp := 3

  // res variable are the test values we're passing to the cound function in main.go
  res := count(b, true)

  if res != exp {
    t.Errorf("Expected %d, got %d instead.\n", exp, res)
  }
}
