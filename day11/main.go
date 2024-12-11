package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	stones := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solve(stones, 25))
	fmt.Println("Part 2: ", solve(stones, 75))
}

type Stone struct {
	num, blinksLeft int
}
func solve(stones []int, blinks int) int {

	cache := make(map[Stone]int)

	count := 0
	for _, s := range stones {
		count += amount(Stone{s, blinks}, &cache)
	}
	return count
}

func amount(s Stone, cache *map[Stone]int) int {
	// Memoization
	if v, ok := (*cache)[s]; ok {
		return v
	}

	// Base case
	if s.blinksLeft==0 {
		return 1
	}

	// Rules, update the cache with recursive calls
	blinks := s.blinksLeft-1
	if s.num == 0 {
		(*cache)[s] = amount(Stone{1, blinks}, cache)
	} else if str := strconv.Itoa(s.num); len(str)%2 == 0 {
		n1 := c.MustAtoi(str[:len(str)/2])
		n2 := c.MustAtoi(str[len(str)/2:])
		(*cache)[s] = amount(Stone{n1, blinks}, cache) + amount(Stone{n2, blinks}, cache)
	} else {
		(*cache)[s] = amount(Stone{s.num * 2024, blinks}, cache)
	}

	return (*cache)[s]
}

func parseInput(path string) []int {
	str := c.Must(os.ReadFile(path))

	re := regexp.MustCompile(`\d+`)

	return c.Map(
		re.FindAllString(string(str), -1),
		c.MustAtoi,
	)
}
