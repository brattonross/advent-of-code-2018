package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type entry struct {
	guardID   int
	entryType entryType
	timestamp time.Time
}

type entryType int

const (
	shiftStart entryType = iota
	fallsAsleep
	wakesUp
)

func (e entry) String() string {
	date := e.timestamp.Format("01/02 15:04")
	switch e.entryType {
	case shiftStart:
		return fmt.Sprintf("[%s] Guard #%d starts shift", date, e.guardID)
	case fallsAsleep:
		return fmt.Sprintf("[%s] Guard #%d falls asleep", date, e.guardID)
	case wakesUp:
		return fmt.Sprintf("[%s] Guard #%d wakes up", date, e.guardID)
	}
	return fmt.Sprintf("unknown event type %#v", e)
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	sort.Strings(lines)

	entries := []entry{}
	for _, l := range lines {
		dateEndIndex := strings.Index(l, "]")
		dateString := l[1:dateEndIndex]
		date, err := time.Parse("2006-01-02 15:04", dateString)
		if err != nil {
			log.Fatalf("failed to parse date %q: %v", dateString, err)
		}
		entry := entry{timestamp: date}
		split := strings.Fields(l[dateEndIndex+2:])
		switch split[0] {
		case "Guard":
			id, err := strconv.Atoi(split[1][1:])
			if err != nil {
				log.Fatalf("failed to parse id %q: %v", split[1][1:], err)
			}
			entry.guardID = id
			entry.entryType = shiftStart
		case "falls":
			entry.guardID = entries[len(entries)-1].guardID
			entry.entryType = fallsAsleep
		case "wakes":
			entry.guardID = entries[len(entries)-1].guardID
			entry.entryType = wakesUp
		}

		entries = append(entries, entry)
	}

	guardSleeps := map[int][]int{}
	for i, e := range entries {
		if e.entryType != wakesUp {
			continue
		}
		if _, ok := guardSleeps[e.guardID]; !ok {
			guardSleeps[e.guardID] = make([]int, 60)
		}
		for min := entries[i-1].timestamp.Minute(); min < e.timestamp.Minute(); min++ {
			guardSleeps[e.guardID][min]++
		}
	}

	sleeper := 0
	sleepMin := 0
	maxCount := 0
	for id, mins := range guardSleeps {
		for min, count := range mins {
			if count > maxCount {
				maxCount = count
				sleepMin = min
				sleeper = id
			}
		}
	}

	log.Printf("Guard #%d, min %d, count %d\n", sleeper, sleepMin, maxCount)
	log.Println(sleeper * sleepMin)
}
