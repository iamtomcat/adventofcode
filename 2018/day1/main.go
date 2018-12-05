package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func checkFrequency(f int, frequencies []int) *int {
	for _, freq := range frequencies {
		if f == freq {
			return &f
		}
	}
	return nil
}

func puzzle1(v []string) (int, int) {
	freq := []int{0}
	sum := 0
	var outFreq *int
	hasChanged := false

	for outFreq == nil {
		for _, value := range v {
			conv, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
			}
			sum += conv

			find := checkFrequency(sum, freq)
			if find == nil {
				freq = append(freq, sum)
			} else if !hasChanged {
				hasChanged = true
				temp := sum
				outFreq = &temp
			}
		}
	}

	return sum, *outFreq
}

func main() {
	dat, err := ioutil.ReadFile("1.txt")

	if err != nil {
		log.Fatal("File doens't exist")
	}

	file := string(dat)
	values := strings.Split(file, "\n")

	sum, freq := puzzle1(values)

	log.Println(sum, freq)
}
