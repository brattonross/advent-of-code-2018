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
		reacted := false
		for i := range units {
			if i == len(units)-1 {
				continue
			}
			if react(units[i], units[i+1]) {
				units = removeUnits(units, i)
				reacted = true
				break
			}
		}
		if !reacted {
			break
		}
	}

	log.Println(len(units))
}

// Determines if two units react together.
func react(a unit, b unit) bool {
	return a.utype == b.utype && a.polarity != b.polarity
}

func removeUnits(u []unit, i int) []unit {
	return append(u[:i], u[i+2:]...)
}
