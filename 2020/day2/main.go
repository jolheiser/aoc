package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(answer1(lines))
	fmt.Println(answer2(lines))
}

func answer1(lines []string) int {
	var num int
	for _, line := range lines {
		s := strings.Split(line, ": ")
		if isValid1(s[0], s[1]) {
			num++
		}
	}
	return num
}

// Should this return a struct? yes
// Does it? no
func parsePolicy(policy string) (int, int, string) {
	p := strings.Split(policy, " ")
	r := strings.Split(p[0], "-")

	char := p[1]
	min, err := strconv.Atoi(r[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(r[1])
	if err != nil {
		panic(err)
	}

	return min, max, char
}

func isValid1(policy, password string) bool {
	min, max, char := parsePolicy(policy)

	var count int
	for _, c := range password {
		if char == string(c) {
			count++
		}
	}

	return count >= min && count <= max
}

func answer2(lines []string) int {
	var num int
	for _, line := range lines {
		s := strings.Split(line, ": ")
		if isValid2(s[0], s[1]) {
			num++
		}
	}
	return num
}

func isValid2(policy, password string) bool {
	pos1, pos2, char := parsePolicy(policy)
	return (string(password[pos1-1]) == char) != (string(password[pos2-1]) == char)
}
