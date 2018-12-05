package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func processLines(l string) []int {
	o := map[string]int{}
	for _, r := range l {
		c := string(r)
		if _, ok := o[c]; ok {
			o[c]++
		} else {
			o[c] = 1
		}
	}

	set := make(map[int]bool)
	for _, v := range o {
		if v == 1 || v == 4 {
			continue
		}
		set[v] = true
	}

	keys := []int{}
	for k := range set {
		keys = append(keys, k)
	}

	return keys
}

func multiplyMap(m map[int]int) int {
	out := 0
	for _, v := range m {
		if out == 0 {
			out = v
		} else {
			out = out * v
		}
	}
	return out
}

func findDiffByOne(s string, value string) string {
	outString := ""
	diffCount := 0
	for idx, r := range s {
		sc := string(r)
		vc := string(value[idx])
		if sc != vc {
			diffCount++
		} else {
			outString += sc
		}
		if diffCount > 1 {
			return ""
		}
	}
	return outString
}

func main() {
	dat, err := ioutil.ReadFile("2.txt")

	if err != nil {
		log.Fatal("File doens't exist")
	}

	file := string(dat)
	values := strings.Split(file, "\n")

	t := map[int]int{}
	for _, l := range values {
		r := processLines(l)
		for _, v := range r {
			if _, ok := t[v]; ok {
				t[v]++
			} else {
				t[v] = 1
			}
		}
	}
	log.Println(multiplyMap(t))

	for idx, v := range values {
		temp := values[idx+1:]
		for _, subv := range temp {
			if diff := findDiffByOne(v, subv); diff != "" {
				println(v, subv, diff)
			}
		}
	}
}
