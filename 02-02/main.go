package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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

	ida, idb := closestStrings(ids)

	log.Printf("%s, %s", ida, idb)

	log.Println(commonChars(ida, idb))
}

// Returns a string which contains the common characters shared by the two strings.
func commonChars(s, t string) string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

// Returns the two strings in the slice which share the highest number of common bytes.
func closestStrings(strings []string) (s string, t string) {
	matched := 0
	for i := 0; i < len(strings)-1; i++ {
		a := strings[i]
		// Compare with all strings after the current one.
		for _, b := range strings[i+1:] {
			match := commonBytes(a, b)
			if match > matched {
				matched = match
				s = a
				t = b
			}
		}
	}
	return s, t
}

// Gets the number of common bytes shared between the two strings.
func commonBytes(s, t string) (match int) {
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			match++
		}
	}
	return match
}
