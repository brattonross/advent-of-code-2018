package main

import (
	"bufio"
	"log"
	"os"
)

// Number of IDs that have exactly two of any letter
// multiplied by number of IDs that have exactly three
// of any letter.
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ids := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		ids = append(ids, s.Text())
	}
	twice, thrice := parseIDs(ids)

	log.Printf("Answer %d", twice*thrice)
}

func parseIDs(ids []string) (twice int, thrice int) {
	for _, id := range ids {
		runes := map[rune]int{}
		for _, r := range id {
			runes[r] = runes[r] + 1
		}
		sawTwice := false
		sawThrice := false
		for _, n := range runes {
			if n == 2 {
				sawTwice = true
			} else if n == 3 {
				sawThrice = true
			}
		}
		if sawTwice {
			twice++
		}
		if sawThrice {
			thrice++
		}
	}
	return twice, thrice
}
