package main

import (
	"fmt"
	"strings"
)

var matrix [][]string

func main() {
	lines := strings.Split(input, "\n")
	matrix = make([][]string, len(lines))
	for idx, line := range lines {
		matrix[idx] = make([]string, len(line))
		for idy, char := range line {
			matrix[idx][idy] = string(char)
		}
	}
	original := copyMatrix(matrix)

	var balance int
	for {
		num := round(4, false)
		if balance == num {
			break
		}
		balance = num
	}
	fmt.Println(balance)

	balance = 0
	matrix = original
	for {
		num := round(5, true)
		if balance == num {
			break
		}
		balance = num
	}
	fmt.Println(balance)
}

func round(tolerance int, far bool) int {
	var num int
	mut := copyMatrix(matrix)
	for idx, row := range matrix {
		for idy, col := range row {
			var adj int
			if far {
				adj = countAdjacentFar(idx, idy, matrix)
			} else {
				adj = countAdjacent(idx, idy, matrix)
			}
			if col == "L" && adj == 0 {
				mut[idx][idy] = "#"
				num++
			} else if col == "#" {
				if adj >= tolerance {
					mut[idx][idy] = "L"
					continue
				}
				num++
			}
		}
	}
	matrix = mut
	//printLayout(matrix)
	return num
}

func countAdjacent(x, y int, matrix [][]string) int {
	var num int
	for i := -1; i <= 1; i++ {
		row := x + i
		if row < 0 || row > len(matrix)-1 {
			continue
		}
		for j := -1; j <= 1; j++ {
			col := y + j
			if col < 0 || col > len(matrix[row])-1 || (row == x && col == y) {
				continue
			}
			if matrix[row][col] == "#" {
				num++
			}
		}
	}
	return num
}

func countAdjacentFar(x, y int, matrix [][]string) int {
	var num int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if checkAdjacentFar(x, y, i, j, matrix) {
				num++
			}
		}
	}
	return num
}

func checkAdjacentFar(x, y, stepX, stepY int, matrix [][]string) bool {
	for {
		x += stepX
		if x < 0 || x > len(matrix)-1 {
			break
		}
		y += stepY
		if y < 0 || y > len(matrix[x])-1 {
			break
		}
		if matrix[x][y] != "." {
			return matrix[x][y] == "#"
		}
	}
	return false
}

func copyMatrix(src [][]string) [][]string {
	c := make([][]string, len(src))
	copy(c, src)
	for idx, cc := range c {
		ccc := make([]string, len(cc))
		copy(ccc, cc)
		c[idx] = ccc
	}
	return c
}

func printLayout(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}
