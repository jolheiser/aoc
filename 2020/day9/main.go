package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	inputs := make([]int, len(lines))
	for idx, line := range lines {
		inputs[idx] = mustInt(line)
	}
	fmt.Println(answer1(25, 25, inputs))
	fmt.Println(answer2(inputs))
}

func answer1(preamble, consider int, inputs []int) int {
	current := inputs[:preamble]

	for _, input := range inputs[preamble:] {
		if !canSum(input, current[len(current)-consider:]) {
			return input
		}
		current = append(current[1:], input)
	}
	return -1
}

func canSum(target int, available []int) bool {
	for idx, num1 := range available {
		for _, num2 := range available[idx+1:] {
			if num1+num2 == target {
				return true
			}
		}
	}
	return false
}

const invalid = 105950735

func answer2(inputs []int) int {
	var low, high, num int
	for idx, input1 := range inputs {
		low, high = math.MaxInt32, 0
		num = input1
		if input1 < low {
			low = input1
		}
		if input1 > high {
			high = input1
		}
		for idy, input2 := range inputs[idx+1:] {
			num += input2
			if input2 < low {
				low = input2
			}
			if input2 > high {
				high = input2
			}
			if num == invalid && idy >= 1 {
				return low + high
			}
			if num > invalid {
				break
			}
		}
	}
	return -1
}

func mustInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
