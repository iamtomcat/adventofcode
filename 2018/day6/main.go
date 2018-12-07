package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func readFile(f string) []string {
	dat, err := ioutil.ReadFile(f)

	if err != nil {
		log.Fatal("File doens't exist")
	}

	file := strings.TrimSpace(string(dat))
	return strings.Split(file, "\n")
}

type position interface {
	xValue() int
	yValue() int
}

type generator struct {
	id         int
	x          int
	y          int
	boundIDs   map[int]bool
	isInfinite bool
}

func (g generator) distance(other position) int {
	xFloat := float64(g.x - other.xValue())
	yFloat := float64(g.y - other.yValue())
	return int(math.Abs(xFloat) + math.Abs(yFloat))
}

type point struct {
	x             int
	y             int
	ownedBy       int
	totalDistance int
}

func (p point) xValue() int {
	return p.x
}

func (p point) yValue() int {
	return p.y
}

func main() {
	f := readFile("6")

	var points = map[int]generator{}

	gX := 0
	gY := 0
	for idx, l := range f {
		s := strings.Split(l, ", ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		p := generator{id: idx, x: x, y: y, boundIDs: map[int]bool{}, isInfinite: false}
		points[p.id] = p
		if p.x > gX {
			gX = p.x
		}
		if p.y > gY {
			gY = p.y
		}
	}

	maxDistanceValue := 10000

	r := int(math.Max(float64(gX), float64(gY))) + 1
	closest := map[int]int{}

	pointsLessThan := []point{}

	for x := 0; x < r; x++ {
		for y := 0; y < r; y++ {

			pointToTest := point{x: x, y: y, ownedBy: -2}

			closestDistance := math.MaxUint8
			for _, p := range points {
				d := p.distance(pointToTest)

				pointToTest.totalDistance += d

				if d < closestDistance {
					pointToTest.ownedBy = p.id
					closestDistance = d
				} else if d == closestDistance {
					pointToTest.ownedBy = -1
				}
			}

			if pointToTest.totalDistance < maxDistanceValue {
				pointsLessThan = append(pointsLessThan, pointToTest)
			}

			if x == 0 || y == 0 || x == r-1 || y == r-1 {
				if pointToTest.ownedBy > -1 {
					p := points[pointToTest.ownedBy]
					p.isInfinite = true
					points[pointToTest.ownedBy] = p
				}
			}

			if _, ok := closest[pointToTest.ownedBy]; !ok {
				closest[pointToTest.ownedBy] = 1
			} else {
				closest[pointToTest.ownedBy]++
			}
		}
	}

	highid := 0
	high := 0
	for _, p := range points {
		if !p.isInfinite {
			if closest[p.id] > high {
				high = closest[p.id]
				highid = p.id
			}
		}
	}

	p := points[highid]
	log.Println(p.id, " ", p.boundIDs, " ", closest[p.id], p.isInfinite)

	log.Println("Points less than ", maxDistanceValue, " are: ", len(pointsLessThan))
}
