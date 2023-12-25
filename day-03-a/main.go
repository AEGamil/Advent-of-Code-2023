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
	start, end := 0, 0
	for lk, line := range lines {
		for j := 0; j < len(line); j++ {
			r := rune(line[j])
			if unicode.IsNumber(r) {
				start = j
				end = start
				for unicode.IsNumber(rune(line[end])) {
					end++
				}
				if valid(lines, lk, start, end) {
					total += parseTheNumber(line[start:end])
				}
				j = end
			}
		}
	}
	fmt.Printf("total: %v\n", total)
}

func parseTheNumber(line string) int {
	i, _ := strconv.Atoi(line)
	return i
}

// checks for all the adjacent cells around the number (including the number xD)
func valid(lines []string, l int, s int, e int) bool {
	bUp, bDw := 1, 2
	if s > 0 {
		s--
	}
	if e < len(lines[0])-1 {
		e++
	}
	if l == 0 {
		bUp = 0
	} else if l == len(lines)-1 {
		bDw = 1
	}
	for i := l - bUp; i < l+bDw; i++ {
		for j := s; j < e; j++ {
			r := rune(lines[i][j])
			if !unicode.IsNumber(r) && r != '.' {
				return true
			}
		}
	}
	return false
}
