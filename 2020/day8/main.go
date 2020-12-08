package main

import (
	"fmt"
	"go.jolheiser.com/set"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(run(lines))

	for idx, line := range lines {
		cp := make([]string, len(lines))
		copy(cp, lines)
		switch line[:3] {
		case "jmp":
			cp[idx] = strings.ReplaceAll(line, "jmp", "nop")
		case "nop":
			cp[idx] = strings.ReplaceAll(line, "nop", "jmp")
		}
		if acc, pass := run(cp); pass {
			fmt.Println(acc)
			break
		}
	}
}

func run(lines []string) (int, bool) {
	var acc, idx int
	ran := set.NewIntSet()
	for {
		if ran.Has(idx) {
			break
		}
		ran.Add(idx)

		instruction := lines[idx]
		parts := strings.Fields(instruction)

		op := parts[0]
		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch op {
		case "acc":
			acc += arg
			idx++
		case "jmp":
			idx += arg
		case "nop":
			idx++
		}

		if idx == len(lines) {
			return acc, true
		}
	}
	return acc, false
}
