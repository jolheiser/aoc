package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")

	fmt.Println(answer1(lines))
	fmt.Println(answer2(lines))
}

func answer1(lines []string) int {
	s := &ship{
		direction: "e",
	}

	for _, line := range lines {
		ins, val := line[:1], line[1:]
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		switch strings.ToLower(ins) {
		case "n":
			s.ns += v
		case "s":
			s.ns -= v
		case "e":
			s.ew += v
		case "w":
			s.ew -= v
		case "f":
			s.forward(v)
		case "r":
			s.right(v)
		case "l":
			s.left(v)
		}
	}

	return s.manhattan()
}

func answer2(lines []string) int {
	s := &ship{
		direction: "e",
		wp: &waypoint{
			ns: 1,
			ew: 10,
		},
	}

	for _, line := range lines {
		ins, val := line[:1], line[1:]
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		switch strings.ToLower(ins) {
		case "n":
			s.wp.ns += v
		case "s":
			s.wp.ns -= v
		case "e":
			s.wp.ew += v
		case "w":
			s.wp.ew -= v
		case "f":
			for i := 0; i < v; i++ {
				s.ns += s.wp.ns
				s.ew += s.wp.ew
			}
		case "r":
			s.right2(v)
		case "l":
			s.left2(v)
		}
	}

	return s.manhattan()
}

type ship struct {
	direction string
	ns        int
	ew        int

	wp *waypoint
}

func (s *ship) left(degree int) {
	turns := degree / 90
	for i := 0; i < turns; i++ {
		switch s.direction {
		case "n":
			s.direction = "w"
		case "s":
			s.direction = "e"
		case "e":
			s.direction = "n"
		case "w":
			s.direction = "s"
		}
	}
}

func (s *ship) right(degree int) {
	turns := degree / 90
	for i := 0; i < turns; i++ {
		switch s.direction {
		case "n":
			s.direction = "e"
		case "s":
			s.direction = "w"
		case "e":
			s.direction = "s"
		case "w":
			s.direction = "n"
		}
	}
}

func (s *ship) right2(degree int) {
	turns := degree / 90
	for i := 0; i < turns; i++ {
		s.wp.ns, s.wp.ew = -s.wp.ew, s.wp.ns
	}
}

func (s *ship) left2(degree int) {
	turns := degree / 90
	for i := 0; i < turns; i++ {
		s.wp.ns, s.wp.ew = s.wp.ew, -s.wp.ns
	}
}

func (s *ship) forward(value int) {
	switch s.direction {
	case "n":
		s.ns += value
	case "s":
		s.ns -= value
	case "e":
		s.ew += value
	case "w":
		s.ew -= value
	}
}

func (s *ship) manhattan() int {
	return int(math.Abs(float64(s.ns)) + math.Abs(float64(s.ew)))
}

type waypoint struct {
	ns int
	ew int
}
