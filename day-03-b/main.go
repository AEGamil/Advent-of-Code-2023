package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// YEEH we don check for errors huh
	dat, _ := os.ReadFile("input.txt")
	inp := string(dat)
	sum(inp)
}

func sum(inp string) {
	total := 0
	lines := strings.Split(inp, "\n")
	for lk, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == '*' {
				total += around_mul(lines, lk, j)
			}
		}
	}
	fmt.Printf("total: %v\n", total)
}
func around_mul(lines []string, l int, p int) int {
	n1, n2, skip := 0, 0, 0
	bU, bD, bR, bL := 1, 2, 2, 1
	if l == 0 {
		bU = 0
	} else if l == len(lines)-1 {
		bD = 1
	}
	if p == 0 {
		bL = 0
	} else if p == len(lines[0])-1 {
		bR = 1
	}
	for i := l - bU; i < l+bD; i++ {
		for j := p - bL; j < p+bR; j++ {
			if unicode.IsNumber(rune(lines[i][j])) {
				if n1 == 0 {
					n1, skip = extractNumber(lines[i], j)
				} else {
					n2, skip = extractNumber(lines[i], j)
				}
				j += skip
			}
		}
	}
	return n1 * n2
}

func extractNumber(line string, p int) (int, int) {
	s, e := p, p
	skip := 0
	for s > 0 {
		s--
		if !unicode.IsNumber(rune(line[s])) {
			s++
			break
		}
	}
	for e < len(line)-1 {
		if unicode.IsNumber(rune(line[e])) {
			e++
			skip++
		} else {
			break
		}
	}
	res, _ := strconv.Atoi(line[s:e])
	return res, skip
}
