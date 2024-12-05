package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ManuelGarciaF/AoC-2024/commons"
)

type Rule struct {
	before, after int
}

func main() {
	rules, pages := parseInput(os.Args[1])

	rulesMap := make(map[int]commons.Set[int])
	for _, r := range rules {
		if _, ok := rulesMap[r.before]; !ok {
			rulesMap[r.before] = commons.NewSet[int]()
		}
		rulesMap[r.before].Add(r.after)
	}

	fmt.Println("Part 1: ", solvePart1(rulesMap, pages))
	fmt.Println("Part 2: ", solvePart2(rulesMap, pages))
}

func solvePart1(rules map[int]commons.Set[int], pages [][]int) int {
	sum := 0
	for _, row := range pages {
		if isValidRow(row, rules) {
			sum += row[len(row)/2]
		}
	}

	return sum
}

func isValidRow(row []int, rules map[int]commons.Set[int]) bool {
	for i := range row {
		if ok, _ := isValidElem(row, i, rules); !ok {
			return false
		}
	}

	return true
}

// Returns if it's valid, and the index of the element that breaks the rule
func isValidElem(row []int, i int, rules map[int]commons.Set[int]) (bool, int) {
	// Check that all previous elements in the row aren't in the after set
	// of the current element
	for _, prev := range row[:i] {
		if rules[row[i]].Contains(prev) {
			return false, i
		}
	}
	return true, -1
}

func solvePart2(rules map[int]commons.Set[int], pages [][]int) int {
	invalid := slices.DeleteFunc(slices.Clone(pages), func(row []int) bool {
		return isValidRow(row, rules)
	})

	sum := 0
	for _, row := range invalid {
		row := fixOrder(row, rules)
		sum += row[len(row)/2]
	}

	return sum
}

func fixOrder(row []int, rules map[int]commons.Set[int]) []int {
	// Fix order left to right
	for !isValidRow(row, rules) {
		for i := range row {
			valid, invalidIndex := isValidElem(row, i, rules)
			if valid {
				continue
			}

			e := row[i]
			row = slices.Delete(row, i, i+1)
			row = slices.Insert(row, invalidIndex-1, e)
		}
	}
	return row
}

func parseInput(path string) ([]Rule, [][]int) {
	file := commons.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make([]Rule, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		nums := strings.Split(scanner.Text(), "|")
		rules = append(rules, Rule{
			before: commons.MustAtoi(nums[0]),
			after:  commons.MustAtoi(nums[1]),
		})

	}

	pages := make([][]int, 0)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		row := make([]int, 0, len(nums))
		for _, s := range nums {
			row = append(row, commons.MustAtoi(s))
		}
		pages = append(pages, row)
	}

	return rules, pages
}
