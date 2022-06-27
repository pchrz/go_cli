package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"flag"
	)

	//bufio -- reads text
	//fmt -- print formatted output
	//io -- provides the io.reader interface
	//os -- so you can use OS resources

func main() { 
  // Defining a boolean flag -l to count the lines instead of words.
  lines := flag.Bool("l", false, "Count Lines")
  // Parsing the flags provided by user
  flag.Parse()

  // Calling the count function to count the number of words
  // received from the Standard Input and printing it out
  fmt.Println(count(os.Stdin, *lines))

}

// For now think of IO.reader as any GO type from which you can read data.

func count(r io.Reader, countLines bool) int {
  // A scanner is used to read text from a Reader (such as files) 
  scanner := bufio.NewScanner(r)

  // If the count lines flag is not set, we want to count words.
  // Define the scanner split type to words (default is split by lines)
  if !countLines {
  scanner.Split(bufio.ScanWords)
  }

  // Defining a counter, this variable will keep track of the words found.
  wc := 0

  // For every word scanned, increment the counter.
  for scanner.Scan() {
	wc++
  }

  // Return the total, funcitons must return something.
  return wc
}

