package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum, err := sum(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(sum)
}

// Sums the digits (line-separated) from the reader.
func sum(r io.Reader) (int, error) {
	sum := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			return 0, err
		}
		sum += n
	}
	return sum, nil
}
