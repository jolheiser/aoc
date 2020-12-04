package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")

	passports := make([]passport, 0)
	var p passport
	for idx, line := range lines {
		if line == "" {
			passports = append(passports, p)
			p = passport{}
		}
		parts := strings.Split(line, " ")
		for _, part := range parts {
			kv := strings.Split(part, ":")
			switch kv[0] {
			case "byr":
				p.byr = kv[1]
			case "iyr":
				p.iyr = kv[1]
			case "eyr":
				p.eyr = kv[1]
			case "hgt":
				p.hgt = kv[1]
			case "hcl":
				p.hcl = kv[1]
			case "ecl":
				p.ecl = kv[1]
			case "pid":
				p.pid = kv[1]
			case "cid":
				p.cid = kv[1]
			}
		}
		if idx == len(lines)-1 {
			passports = append(passports, p)
		}
	}

	var num1, num2 int
	for _, p := range passports {
		if p.valid1() {
			num1++
		}
		if p.valid2() {
			num2++
		}
	}
	fmt.Println(num1)
	fmt.Println(num2)
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) valid1() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

var hclre = regexp.MustCompile(`#[0-9a-f]{6}`)

func (p passport) valid2() bool {
	b, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if b < 1920 || b > 2002 {
		return false
	}

	i, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if i < 2010 || i > 2020 {
		return false
	}

	e, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if e < 2020 || e > 2030 {
		return false
	}

	if strings.HasSuffix(p.hgt, "cm") {
		h, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm"))
		if err != nil {
			return false
		}
		if h < 150 || h > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		h, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in"))
		if err != nil {
			return false
		}
		if h < 59 || h > 76 {
			return false
		}
	} else {
		return false
	}

	if !hclre.MatchString(p.hcl) {
		return false
	}

	switch p.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	_, err = strconv.Atoi(p.pid)
	if err != nil {
		return false
	}

	return len(p.pid) == 9
}
