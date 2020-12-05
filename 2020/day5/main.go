package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	seats := make([]seat, len(lines))
	for idx, line := range lines {
		seats[idx] = decode(line)
	}

	seatMap := make(map[int]struct{})
	var high int
	for _, s := range seats {
		seatMap[s.id()] = struct{}{}
		if s.id() > high {
			high = s.id()
		}
	}
	fmt.Println(high)

	var seatID int
	for id := range seatMap {
		_, up := seatMap[id+1]
		if !up {
			_, up = seatMap[id+2]
			if up {
				seatID = id + 1
				break
			}
		}
		_, down := seatMap[id-1]
		if !down {
			_, down = seatMap[id-2]
			if down {
				seatID = id - 1
				break
			}
		}
	}
	fmt.Println(seatID)
}

type seat struct {
	row int
	col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func decode(key string) seat {
	rowKey, colKey := key[0:7], key[7:10]

	parse := func(k rune, low, high int) (int, int) {
		mid := (low + high) / 2
		switch k {
		case 'F', 'L':
			return low, mid
		case 'B', 'R':
			return mid + 1, high
		}
		return 0, 0
	}

	low, high := 0, 127
	for _, r := range rowKey {
		low, high = parse(r, low, high)
	}
	row := low

	low, high = 0, 7
	for _, c := range colKey {
		low, high = parse(c, low, high)
	}
	col := low

	return seat{
		row: row,
		col: col,
	}
}
