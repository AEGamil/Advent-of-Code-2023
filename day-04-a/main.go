package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// YEEH we don check for errors huh
	dat, _ := os.ReadFile("input.txt")
	inp := string(dat)
	sum(inp)
}

// 2^(n-1), where n is the number of matchs
func sum(inp string) {
	total := 0
	lines := strings.Split(inp, "\n")
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		// 0: the winning list , 1: my number picks
		card := strings.Split(line, "|")
		if p := compare(parse_nums(card[0]), parse_nums(card[1])); p > 0 {
			total += int(math.Pow(float64(2), float64(p-1)))
		}
	}
	fmt.Printf("total: %v\n", total)
}

// this function doesn't work well if the input.txt is CRLF, so be a good boy and make it LF :)
func parse_nums(list string) []int {
	nums := []int{}
	e := strings.Split(list, " ")
	// fmt.Printf("e: %v\n", e)
	e = slices.DeleteFunc(e, func(s string) bool { return s == "" })
	for _, s := range e {
		if num, err := strconv.Atoi(s); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func compare(w []int, m []int) int {
	matches := 0
	for _, e := range w {
		if slices.Contains(m, e) {
			matches++
		}
	}
	return matches
}
