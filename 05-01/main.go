package main

import (
	"io/ioutil"
	"log"
)

type unit struct {
	utype    byte
	polarity polarity
}

func (u unit) String() string {
	if u.polarity == upper {
		return string(u.utype - 0x20)
	}
	return string(u.utype)
}

type polarity int

const (
	lower polarity = iota
	upper
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	units := []unit{}
	for _, u := range b {
		p := lower
		if u < 'a' {
			u += 0x20
			p = upper
		}
		units = append(units, unit{u, p})
	}

	for {
		removed := 0
		for i := range units {
			j := i - removed
			if j == len(units)-1 {
				continue
			}
			if react(units[j], units[j+1]) {
				units = removeUnits(units, j)
				removed += 2
			}
		}
		if removed == 0 {
			break
		}
	}

	log.Println(units)
}

// Determines if two units react together.
func react(a unit, b unit) bool {
	return a.utype == b.utype && a.polarity != b.polarity
}

func removeUnits(u []unit, i int) []unit {
	return append(u[:i], u[i+2:]...)
}
