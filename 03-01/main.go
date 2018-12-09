package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	claims := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		claims = append(claims, s.Text())
	}

	fabric := parseClaims(claims)
	count := countOverlaps(fabric)

	log.Print(count)
}

func countOverlaps(fabric map[int]map[int]int) (count int) {
	for _, x := range fabric {
		for _, y := range x {
			if y > 1 {
				count++
			}
		}
	}
	return count
}

func parseClaims(claims []string) map[int]map[int]int {
	fabric := map[int]map[int]int{}
	for _, c := range claims {
		var id, x, y, w, h int

		fmt.Sscanf(c, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)

		for n := x; n < x+w; n++ {
			if _, ok := fabric[n]; !ok {
				fabric[n] = map[int]int{}
			}
			for o := y; o < y+h; o++ {
				fabric[n][o]++
			}
		}
	}
	return fabric
}
