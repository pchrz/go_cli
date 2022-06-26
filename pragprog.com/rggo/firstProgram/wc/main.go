package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	)

	//bufio -- reads text
	//fmt -- print formatted output
	//io -- provides the io.reader interface
	//os -- so you can use OS resources

func main() { 
  // Calling the count function to count the number of words
  // received from the Standard Input and printing it out
  fmt.Println(count(os.Stdin))

}

// For now think of IO.reader as any GO type from which you can read data.

func count(r io.Reader) int {
  // A scanner is used to read text from a Reader (such as files) 
  scanner := bufio.NewScanner(r)

  // Define the scanner split type to words (default is split by lines)
  scanner.Split(bufio.ScanWords)

  // Defining a counter, this variable will keep track of the words found.
  wc := o

  // For every word scanned, increment the counter.
  for scanner.Scan() {
	wc++
  }

  // Return the total, funcitons must return something.
  return wc
}

