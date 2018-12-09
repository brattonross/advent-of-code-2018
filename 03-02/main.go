package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct{ id, x, y, w, h int }

func newClaim(s string) *claim {
	var id, x, y, w, h int
	fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
	return &claim{id, x, y, w, h}
}

func claimsToMap(claims []*claim) map[int]map[int]int {
	fabric := map[int]map[int]int{}
	for _, c := range claims {
		for n := c.x; n < c.x+c.w; n++ {
			if _, ok := fabric[n]; !ok {
				fabric[n] = map[int]int{}
			}
			for m := c.y; m < c.y+c.h; m++ {
				fabric[n][m]++
			}
		}
	}
	return fabric
}

func main() {
	// 1 1 1 1
	// 1 1 2 2
	// 1 1 2 3
	// 1 1 2 2

	// Parse claims
	// Map to 2D map of number of overlaps
	// For each claim, check the map to see if all positions that it occupies are 1

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	claims := []*claim{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		claims = append(claims, newClaim(s.Text()))
	}

	fabric := claimsToMap(claims)

	for _, c := range claims {
		ok := true

		for n := c.x; n < c.x+c.w; n++ {
			for o := c.y; o < c.y+c.h; o++ {
				if fabric[n][o] != 1 {
					ok = false
					break
				}
			}

			if !ok {
				break
			}
		}

		if ok {
			log.Printf("found %d", c.id)
		}
	}
}
