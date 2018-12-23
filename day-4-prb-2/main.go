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

type xy struct {
	x int
	y []int
}
type event struct {
	id   int
	kind eventKind
	time time.Time
}

type eventKind byte

const (
	eventStart = iota
	eventAsleep
	eventAwake
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	sort.Strings(lines)

	var events []event

	for _, line := range lines {
		dateStart := strings.Index(line, "]")
		dateText := line[1:dateStart]
		date, err := time.Parse("2006-01-02 15:04", dateText)
		if err != nil {
			log.Fatalf("could not parser date %q: %v", dateText, err)
		}
		e := event{time: date}
		pieces := strings.Fields(line[dateStart+2:])
		switch pieces[0] {
		case "Guard":
			id, err := strconv.Atoi(pieces[1][1:])
			if err != nil {
				log.Fatalf("could not parse id %q: %v", pieces[1][1:], err)
			}
			e.id = id
			e.kind = eventStart
		case "falls":
			e.id = events[len(events)-1].id
			e.kind = eventAsleep

		case "wakes":
			e.id = events[len(events)-1].id
			e.kind = eventAwake
		}

		events = append(events, e)
	}

	guardID, minute := findGuard(events)
	fmt.Println(guardID * minute)

}

func findGuard(events []event) (id, minute int) {

	minutes := make([]map[int]int, 60)
	for i := range minutes {
		minutes[i] = make(map[int]int)
	}

	for i, e := range events {
		if e.kind != eventAwake {
			continue
		}
		for i := events[i-1].time.Minute(); i < e.time.Minute(); i++ {
			minutes[i][e.id]++
		}
	}

	maxMinute := 0
	maxID := 0
	maxCount := 0
	for minute, counts := range minutes {
		for id, n := range counts {
			if n > maxCount {
				maxMinute = minute
				maxCount = n
				maxID = id
			}
		}
	}
	return maxID, maxMinute
}
