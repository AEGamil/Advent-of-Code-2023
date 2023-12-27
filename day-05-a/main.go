package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// YEEH we don check for errors huh
	dat, _ := os.ReadFile("input.txt")
	inp := string(dat)
	// in this section i'm modifying the input as im too lazy to parse the seeds alone xD
	x := []byte(inp)
	x[6] = '\n'
	inp = string(x)
	//
	sol(inp)
}

func sol(inp string) {
	result := 0
	m := strings.Split(inp, "\n\n")
	seeds := Str2Int(m[0])[0]
	maps := m[1:]
	for _, seed := range seeds {
		for _, maap := range maps {
			// n_map is the values in the maps [][]int
			n_map := Str2Int(maap)
			seed = mapvalue(n_map, seed)
		}
		if result == 0 || seed < result {
			result = seed
		}
	}
	fmt.Printf("result: %v\n", result)
}

// this will seperate the map names and the numbers
func Str2Int(x string) [][]int {
	temp := [][]int{}
	s := strings.Split(x, ":")
	s = strings.Split(s[1], "\n")[1:]
	for i := 0; i < len(s); i++ {
		t := strings.Split(s[i], " ")
		temp = append(temp, []int{})
		for j := 0; j < len(t); j++ {
			if num, err := strconv.Atoi(t[j]); err == nil {
				temp[i] = append(temp[i], num)
			}
		}
	}
	return temp
}

func mapvalue(maap [][]int, value int) int {
	for i := 0; i < len(maap); i++ {
		if value >= maap[i][1] && value <= maap[i][1]+maap[i][2] {
			return maap[i][0] + value - maap[i][1]
		}
	}
	return value
}
