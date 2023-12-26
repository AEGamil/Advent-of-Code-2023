package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// YEEH we don check for errors huh
	dat, _ := os.ReadFile("input.txt")
	inp := string(dat)
	// result is 6857330 (gonna provide the result in a better way later i guess)
	sum(inp)
}

func sum(inp string) {
	total, temp, p := 0, 0, 0
	queue := []int{1}
	lines := strings.Split(inp, "\n")
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		// 0: the winning list , 1: my number picks
		card := strings.Split(line, "|")
		p = compare(parse_nums(card[0]), parse_nums(card[1]))
		if len(queue) > 0 {
			// DEQUEUING
			// temp is the number of copies of the current card
			temp += queue[0]
			queue = queue[1:]
		}

		if p > len(queue) {
			queue = append(queue, make([]int, p-len(queue))...)
		}
		for i := 0; i < p; i++ {
			queue[i] += temp
		}
		total += temp
		temp, p = 1, 0
	}
	fmt.Printf("total: %v\n", total)

}

// this function doesn't work well if the input.txt is CRLF, so be a good boy and make it LF :)
func parse_nums(list string) []int {
	nums := []int{}
	e := strings.Split(list, " ")
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
