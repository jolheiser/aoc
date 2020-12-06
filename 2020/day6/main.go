package main

import (
	"fmt"
	"strings"
)

func main() {
	groups := strings.Split(input, "\n\n")

	var sum1, sum2 int
	for _, group := range groups {
		sum1 += anyYes(group)
		sum2 += allYes(group)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func anyYes(group string) int {
	lines := strings.Split(group, "\n")

	m := make(map[rune]struct{})
	for _, line := range lines {
		for _, r := range line {
			if _, ok := m[r]; !ok {
				m[r] = struct{}{}
			}
		}
	}
	return len(m)
}

func allYes(group string) int {
	lines := strings.Split(group, "\n")

	m := make(map[rune]int)
	for _, line := range lines {
		for _, r := range line {
			_, ok := m[r]
			if ok {
				m[r] += 1
				continue
			}
			m[r] = 1
		}
	}

	var sum int
	for _, num := range m {
		if num == len(lines) {
			sum++
		}
	}

	return sum
}
