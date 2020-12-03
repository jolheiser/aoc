package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	v := make([][]string, len(lines))
	for idx, line := range lines {
		l := make([]string, len(line))
		for idy, c := range line {
			l[idy] = string(c)
		}
		v[idx] = l
	}

	m := matrix{m: v}
	fmt.Println(m.slope(3, 1))
	fmt.Println(m.slope(1, 1) *
		m.slope(3, 1) *
		m.slope(5, 1) *
		m.slope(7, 1) *
		m.slope(1, 2))
}

type matrix struct {
	m [][]string
}

func (m matrix) height() int {
	return len(m.m)
}

func (m matrix) width() int {
	return len(m.m[0])
}

func (m matrix) slope(right, down int) int {
	width, height := m.width(), m.height()

	row, col, trees := 0, 0, 0
	for {
		if row >= height {
			break
		}
		if col >= width {
			col -= width
		}

		if m.m[row][col] == "#" {
			trees++
		}

		row += down
		col += right
	}
	return trees
}
