package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func readFile(f string) []string {
	dat, err := ioutil.ReadFile(f)

	if err != nil {
		log.Fatal("File doens't exist")
	}

	file := string(dat)
	return strings.Split(file, "\n")
}

func removePairs(line string) string {
	if len(line) <= 1 {
		return line
	}

	first := rune(line[0])
	second := rune(line[1])

	if first != second && strings.ToLower(string(line[0])) == strings.ToLower(string(line[1])) {
		// log.Println(line[:2], " ", line[2:])
		return removePairs(line[2:])
	}

	remStr := string(line[0]) + removePairs(line[1:])

	// remStr := removePairsWithRemove(line[1:], string(line[0]))

	// log.Println("remstr", remStr, lastRemoved)

	if len(remStr) > 1 {
		firstRem := rune(remStr[0])
		secondRem := rune(remStr[1])

		if firstRem != secondRem && strings.ToLower(string(remStr[0])) == strings.ToLower(string(remStr[1])) {
			// log.Println("2", remStr[:2], " ", remStr[2:])
			return removePairs(remStr[2:])
		}
	}

	return remStr
}

func main() {
	f := readFile("5")

	wg := new(sync.WaitGroup)

	for _, l := range f {
		// data := removePairs(l)
		// log.Println(len(data))

		chars := "abcdefghijklmnopqrstuvwxyz"

		// lowest := 0
		// letter := ""
		for _, c := range chars {
			sc := string(c)
			s := strings.Replace(l, sc, "", -1)
			s2 := strings.Replace(s, strings.ToUpper(sc), "", -1)

			wg.Add(1)
			go func() {
				data := removePairs(s2)
				count := len(data)
				log.Println(count, " ", sc)
				wg.Done()
			}()

			// if lowest > count || lowest == 0 {
			// 	lowest = count
			// 	letter = sc
			// }
		}

	}
	wg.Wait()
}
