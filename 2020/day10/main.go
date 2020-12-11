package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	inputs := make([]int, len(lines))
	for idx, line := range lines {
		inputs[idx] = mustInt(line)
	}
	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)

	diffs := make(map[int]int)
	for k, v := range inputs {
		var prev int
		if k != 0 {
			prev = inputs[k-1]
		}
		diffs[v-prev]++
	}
	fmt.Println(answer1(inputs))
	fmt.Println(answer2(inputs))
}

func answer1(inputs []int) int {
	diffs := make(map[int]int)
	for k, v := range inputs {
		var prev int
		if k != 0 {
			prev = inputs[k-1]
		}
		diffs[v-prev]++
	}
	return diffs[1] * diffs[3]
}

func answer2(inputs []int) int {
	arrangements := make(map[int]int)
	arrangements[0] = 1

	for _, input := range inputs {
		arrangements[input] = arrangements[input-1] + arrangements[input-2] + arrangements[input-3]
	}
	return arrangements[inputs[len(inputs)-1]]
}

func mustInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
