package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	file := commons.Must(os.Open(os.Args[1]))
	defer file.Close()

	text := string(commons.Must(io.ReadAll(file)))

	fmt.Println("Part 1: ", solvePart1(text))
	fmt.Println("Part 2: ", solvePart2(text))
}

func solvePart1(text string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(text, -1)

	total := 0
	for _, m := range matches {
		total += commons.MustAtoi(m[1]) * commons.MustAtoi(m[2])
	}

	return total
}

func solvePart2(text string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(text, -1)

	enabled := true
	total := 0
	fmt.Println(matches)
	for _, m := range matches {
		switch m[0] {
		case "do()":
			enabled = true
			continue
		case "don't()":
			enabled = false
			continue
		}
		if enabled {
			total += commons.MustAtoi(m[1]) * commons.MustAtoi(m[2])
		}
	}

	return total
}
