package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guard struct {
	id     int
	events []event
}

type guardEvent struct {
	startMinute int
	endMinute   int
}

func (g guard) sumTimeAsleep() int {
	// duration := time.Since(then)
	var timeSum float64
	var startTime time.Time
	for _, e := range g.events {
		if strings.HasPrefix(e.info, "falls") {
			startTime = e.date
		} else if strings.HasPrefix(e.info, "wakes") {
			duration := e.date.Sub(startTime)
			min := duration.Minutes()
			timeSum += min
		}
	}

	return int(timeSum)
}

func (g guard) getStartingMinutes() []guardEvent {
	var out []guardEvent

	var temp guardEvent

	for _, e := range g.events {
		if strings.HasPrefix(e.info, "Guard") {
			continue
		}

		if strings.HasPrefix(e.info, "wakes") {
			temp.endMinute = e.date.Minute()
		} else if strings.HasPrefix(e.info, "falls") {
			temp = guardEvent{
				startMinute: e.date.Minute(),
			}
		}

		if temp.endMinute != 0 {
			out = append(out, temp)
		}
	}
	return out
}

func (g guard) overlappingMinutes() []guardEvent {
	events := g.getStartingMinutes()
	overlappingEvents := []guardEvent{}

	for x := 0; x < len(events); x++ {
		for y := len(events) - 1; y > x; y-- {
			if x == y {
				continue
			}
			if events[x].startMinute <= events[y].endMinute && events[y].startMinute <= events[x].endMinute {
				log.Println("overlap", events[x], events[y])
				overlappingEvents = append(overlappingEvents, events[x], events[y])
			}
		}
	}
	return overlappingEvents
}

func (g guard) minuteMostAsleep() int {
	events := g.getStartingMinutes()

	cheese := map[int]int{}
	for x := 0; x < 60; x++ {
		for _, e := range events {
			if x < e.endMinute && x >= e.startMinute {
				cheese[x]++
			}
		}
	}

	var highestMinute = 0
	var minute = 0
	for key, v := range cheese {
		if v > highestMinute {
			highestMinute = v
			minute = key
		}
	}

	return minute
}

func (g guard) mostFrequentlyAsleep() (int, int) {
	overlap := g.overlappingMinutes()

	cheese := map[int]int{}
	for x := 0; x < 60; x++ {
		for _, e := range overlap {
			if x < e.endMinute && x >= e.startMinute {
				cheese[x]++
			}
		}
	}

	highestMinute := 0
	numberOfTimes := 0
	for key, v := range cheese {
		if v > numberOfTimes {
			numberOfTimes = v
			highestMinute = key
		}

	}

	return highestMinute, numberOfTimes
}

type event struct {
	date           time.Time
	info           string
	startingMinute int
}

func (e event) guardID() int {
	s := strings.Split(e.info, " ")
	id := s[1][1:]
	intid, err := strconv.Atoi(id)

	if err != nil {
		return -1
	}

	return intid
}

func setupEvents(guard []string) []event {
	events := []event{}
	for _, i := range guard {
		t := i[1:17]
		d, err := time.Parse("2006-01-02 15:04", t)

		if err != nil {
			log.Fatal(err)
		}

		e := event{
			date:           d,
			info:           i[19:],
			startingMinute: d.Minute(),
		}
		events = append(events, e)
	}
	return events
}

func readFile(f string) []string {
	dat, err := ioutil.ReadFile(f)

	if err != nil {
		log.Fatal("File doens't exist")
	}

	file := string(dat)
	return strings.Split(file, "\n")
}

func part1(guards map[int]guard) {
	gid := 0
	gsummedMinutes := 0
	for _, g := range guards {
		sum := g.sumTimeAsleep()
		if sum > gsummedMinutes {
			gsummedMinutes = sum
			gid = g.id
		}
	}

	log.Println(gid, guards[gid].sumTimeAsleep())
	log.Println(guards[gid].getStartingMinutes())
	log.Println(guards[gid].overlappingMinutes())
}

func part2(guards map[int]guard) {
	gid := 0
	mostFrequent := 0
	minute := 0
	for _, g := range guards {
		min, mf := g.mostFrequentlyAsleep()
		if mf > mostFrequent {
			mostFrequent = mf
			gid = g.id
			minute = min
		}
	}
	log.Println("guard", gid)
	log.Println(guards[gid].minuteMostAsleep())
	log.Println(minute, mostFrequent)
}

func main() {
	guardActions := readFile("4")

	events := setupEvents(guardActions)

	sort.Slice(events, func(i, j int) bool {
		return events[i].date.Before(events[j].date)
	})

	guards := map[int]guard{}

	log.Println(guards)

	guardID := 0
	for _, e := range events {
		id := e.guardID()

		if id != -1 {
			guardID = id
		}

		if g, ok := guards[guardID]; ok {
			g.events = append(g.events, e)
			guards[guardID] = g
		} else {
			guards[guardID] = guard{
				id:     guardID,
				events: []event{e},
			}
		}
	}

	part1(guards)
	part2(guards)
}
