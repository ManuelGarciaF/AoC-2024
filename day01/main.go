package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	l, r := parseInput(os.Args[1])
	fmt.Println("Part 1: ", solvePart1(l, r))
	fmt.Println("Part 2: ", solvePart2(l, r))
}

func solvePart1(l, r []int) int {
	slices.Sort(l)
	slices.Sort(r)
	if len(l) != len(r) {
		panic("invalid input")
	}
	sum := 0
	for i := range l {
		sum += commons.Abs(l[i] - r[i])
	}

	return sum
}

func solvePart2(l, r []int) int {
	occurrences := make(map[int]int, len(r))
	for _, v := range r {
		occurrences[v] += 1
	}

	sum := 0
	for _, v := range l {
		sum += v * occurrences[v]
	}

	return sum
}

func parseInput(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		ints := strings.Split(scanner.Text(), "   ")
		l, err := strconv.Atoi(ints[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(ints[1])
		if err != nil {
			panic(err)
		}
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
