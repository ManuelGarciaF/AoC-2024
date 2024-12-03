package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	input := parseInput(os.Args[1])
	fmt.Println("Part 1: ", solve(input, false))
	fmt.Println("Part 2: ", solve(input, true))
}

func solve(reports [][]int, removeOne bool) int {
	count := 0
	for _, report := range reports {
		if isSafe(report, removeOne) {
			count++
		}
	}

	return count
}

func isSafe(report []int, removeOne bool) bool {
	increasing := report[1] > report[0]

	for i := 0; i < len(report)-1; i++ {
		notOrdered := increasing && report[i+1] < report[i] ||
			!increasing && report[i+1] > report[i]
		invalidDiff := report[i] == report[i+1] || commons.Abs(report[i]-report[i+1]) > 3

		if notOrdered || invalidDiff {
			return removeOne && tryRemoveOne(report)
		}
	}
	return true
}

// Stupid bruteforce solution
func tryRemoveOne(report []int) bool {
	for i := range report {
		if isSafe(slices.Delete(slices.Clone(report), i, i+1), false) {
			return true
		}
	}
	return false
}

func parseInput(path string) [][]int {
	file := commons.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0)

	for scanner.Scan() {
		intsStrings := strings.Split(scanner.Text(), " ")
		line := make([]int, 0, len(intsStrings))
		for _, intS := range intsStrings {
			line = append(line, commons.MustAtoi(intS))
		}
		reports = append(reports, line)
	}

	return reports
}
