package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type square struct {
	id    string
	start position
	size  size
}

func render(squares []square) string {
	board := [][]string{}
	fabricCount := 0

	untouchedClaims := map[string]bool{}

	for _, s := range squares {
		isUntouched := true
		// log.Println("Adding square", s.id)
		for x := 0; x < s.start.x+s.size.width; x++ {
			temp := []string{}
			if len(board) > x {
				temp = board[x]
			} else {
				board = append(board, []string{})
			}

			for y := 0; y < s.start.y+s.size.height; y++ {
				if s.start.y <= y && y <= s.start.y+s.size.height && s.start.x <= x && x <= s.start.x+s.size.width {
					if len(temp) > y {
						conv, _ := strconv.Atoi(temp[y])
						if conv > 0 {
							delete(untouchedClaims, temp[y])
							fabricCount++
							temp[y] = "X"
							isUntouched = false
						} else if temp[y] == "X" {
							isUntouched = false
							//do nothing
						} else {
							temp[y] = s.id
						}
					} else {
						temp = append(temp, s.id)
					}
				} else {
					if len(temp) > y {
						// temp[y] = "."
					} else {
						temp = append(temp, ".")
					}
				}
			}
			board[x] = temp
		}
		if isUntouched {
			untouchedClaims[s.id] = true
		}
	}

	log.Println("Untouched claim", untouchedClaims)

	out := "\n"
	// for _, x := range board {
	// 	for _, y := range x {
	// 		out += y
	// 	}
	// 	out += "\n"
	// }

	log.Println("fabric count", fabricCount)

	return out
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func toSquare(s string) square {
	split := strings.Split(s, " ")

	square := square{}

	for idx, v := range split {
		if idx == 0 {
			square.id = v[1:]
		} else if idx == 2 {
			pos := strings.Split(v, ",")
			start := position{x: 0, y: 0}
			for idx2, coord := range pos {
				if idx2 == 0 {
					conv, _ := strconv.Atoi(coord)
					start.x = conv
				} else {
					conv, _ := strconv.Atoi(coord[:len(coord)-1])
					start.y = conv
				}
			}
			square.start = start
		} else if idx == 3 {
			sizes := strings.Split(v, "x")

			start := size{width: 0, height: 0}
			for idx2, coord := range sizes {
				conv, _ := strconv.Atoi(coord)
				if idx2 == 0 {
					start.width = conv
				} else {
					start.height = conv
				}
			}
			square.size = start
		}
	}
	return square
}

func main() {
	dat, err := ioutil.ReadFile("3")

	if err != nil {
		log.Fatal("File doens't exist")
	}

	file := string(dat)
	lines := strings.Split(file, "\n")

	squares := []square{}
	for _, v := range lines {
		squares = append(squares, toSquare(v))
	}

	log.Print(render(squares))

}
