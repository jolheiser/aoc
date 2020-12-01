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
	for idx, line1 := range lines {
		for idy, line2 := range lines {
			if idx == idy {
				continue
			}
			num1, err := strconv.Atoi(line1)
			if err != nil {
				panic(err)
			}
			num2, err := strconv.Atoi(line2)
			if err != nil {
				panic(err)
			}

			if num1+num2 == 2020 {
				return num1 * num2
			}
		}
	}
	return 0
}

func answer2(lines []string) int {
	for idx, line1 := range lines {
		for idy, line2 := range lines {
			for idz, line3 := range lines {
				if idx == idy || idx == idz || idy == idz {
					continue
				}
				num1, err := strconv.Atoi(line1)
				if err != nil {
					panic(err)
				}
				num2, err := strconv.Atoi(line2)
				if err != nil {
					panic(err)
				}
				num3, err := strconv.Atoi(line3)
				if err != nil {
					panic(err)
				}

				if num1+num2+num3 == 2020 {
					return num1 * num2 * num3
				}
			}
		}
	}
	return 0
}
