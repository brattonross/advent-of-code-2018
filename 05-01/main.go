package main

import (
	"io/ioutil"
	"log"
)

type unit struct {
	utype    byte
	polarity polarity
	reacted  bool
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
		units = append(units, unit{u, p, false})
	}

	// Remove reacting units
	log.Println(react(units))
}

func react(units []unit) int {
	remaining := 0
	for i := 0; i < len(units); i++ {

	}

	return remaining
}
