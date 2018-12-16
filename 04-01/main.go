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

// Log file entry
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
	// Parse input
	// Sort into chronological order
	// Find guard that sleeps the most amount
	// Find guard ID where they are asleep at same hour and minute on two days

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

	// Sum sleep durations per guard
	sleepDurations := map[int]time.Duration{}
	for i, e := range entries {
		if e.entryType == wakesUp {
			sleepDurations[e.guardID] += e.timestamp.Sub(entries[i-1].timestamp)
		}
	}

	// Find guard that sleeps the most
	sleepiest := 0
	var maxSleep time.Duration
	for id, dur := range sleepDurations {
		if dur > maxSleep {
			maxSleep = dur
			sleepiest = id
		}
	}

	// Find the minute that the guard was most asleep on
	minutes := make([]int, 60)
	for i, e := range entries {
		if e.guardID != sleepiest || e.entryType != wakesUp {
			continue
		}
		for min := entries[i-1].timestamp.Minute(); min < e.timestamp.Minute(); min++ {
			minutes[min]++
		}
	}

	log.Println(minutes)

	maxMinIndex := 0
	maxMins := 0
	for i, m := range minutes {
		if m > maxMins {
			maxMinIndex = i
			maxMins = m
		}
	}

	log.Println(sleepiest, maxMinIndex)
	log.Println(sleepiest * maxMinIndex)
}
