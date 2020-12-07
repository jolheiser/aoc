package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	containsRe = regexp.MustCompile(`(\d)\s(\w+\s\w+) bags?`)
	bags       = make(map[string]map[string]int)
)

func main() {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		color := strings.Join(strings.Fields(line)[:2], " ")
		if _, ok := bags[color]; !ok {
			bags[color] = make(map[string]int)
		}
		for _, match := range containsRe.FindAllStringSubmatch(line, -1) {
			count, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			bags[color][match[2]] = count
		}
	}

	needle := "shiny gold"
	var num int
	for color, _ := range bags {
		if needle == color {
			continue
		}
		if findOne(needle, color) {
			num++
		}
	}
	fmt.Println(num)
	fmt.Println(count(needle))
}

func findOne(needle, haystack string) bool {
	if needle == haystack {
		return true
	}
	for color, _ := range bags[haystack] {
		if findOne(needle, color) {
			return true
		}
	}
	return false
}

func count(needle string) int {
	var num int
	for color, contains := range bags[needle] {
		num += contains + (contains * count(color))
	}
	return num
}
