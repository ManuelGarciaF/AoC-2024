package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

type Equation struct {
	Target uint64
	Nums   []uint64
}

func main() {
	eqs := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solve(eqs, false))
	fmt.Println("Part 2: ", solve(eqs, true))
}

func solve(eqs []Equation, concat bool) uint64 {
	sum := uint64(0)
	for _, eq := range eqs {
		if canBeMadeTrue(eq.Target, eq.Nums[0], eq.Nums[1:], concat) {
			sum += eq.Target
		}
	}
	return sum
}

func canBeMadeTrue(target, curr uint64, left []uint64, concat bool) bool {
	if len(left) > 0 {
		if curr > target { // Early stop, numbers can only increase
			return false
		}

		// Enable or disable concatenation
		if concat {
			next := left[0]

			// Move curr to the right by the amount of digits of the next num and add it.
			numberOfDigits := int(math.Log10(float64(next))) + 1 // Cast to int to round down
			concatenated := curr*uint64(math.Pow10(numberOfDigits)) + next

			if canBeMadeTrue(target, concatenated, left[1:], concat) {
				return true
			}
		}

		return canBeMadeTrue(target, curr+left[0], left[1:], concat) ||
			canBeMadeTrue(target, curr*left[0], left[1:], concat)
	}
	return curr == target
}

func parseInput(path string) []Equation {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re1 := regexp.MustCompile(`(\d+): (.*)`)
	re2 := regexp.MustCompile(`\d+`)

	eqs := make([]Equation, 0)

	for scanner.Scan() {
		r1 := re1.FindStringSubmatch(scanner.Text())
		r2 := re2.FindAllString(r1[2], -1)
		eqs = append(eqs, Equation{
			Target: c.Must(strconv.ParseUint(r1[1], 10, 64)),
			Nums: c.Map(r2, func(s string) uint64 {
				return c.Must(strconv.ParseUint(s, 10, 64))
			}),
		})
	}

	return eqs
}
