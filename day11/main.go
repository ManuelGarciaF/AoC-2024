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

func solve(stones []int, blinks int) int {

	type Stone struct {
		num, blinksLeft int
	}

	var amount func(Stone) int
	amount = c.Memoize(func(s Stone) int {
		// Base case
		if s.blinksLeft == 0 {
			return 1
		}

		// Rules
		blinks := s.blinksLeft - 1
		if s.num == 0 {
			return amount(Stone{1, blinks})
		}

		str := strconv.Itoa(s.num)
		if len(str)%2 == 0 {
			// Split number in halves
			half := len(str)/2
			n1 := c.MustAtoi(str[:half])
			n2 := c.MustAtoi(str[half:])
			return amount(Stone{n1, blinks}) + amount(Stone{n2, blinks})
		}

		return amount(Stone{s.num * 2024, blinks})
	})

	count := 0
	for _, s := range stones {
		count += amount(Stone{s, blinks})
	}
	return count
}

func parseInput(path string) []int {
	str := c.Must(os.ReadFile(path))

	re := regexp.MustCompile(`\d+`)

	return c.Map(
		re.FindAllString(string(str), -1),
		c.MustAtoi,
	)
}
