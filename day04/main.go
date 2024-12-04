package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ManuelGarciaF/AoC-2024/commons"
)

type XMAS int

const (
	X XMAS = iota
	M
	A
	S
)

func main() {
	input := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(input))
	fmt.Println("Part 2: ", solvePart2(input))
}

func solvePart1(letters [][]XMAS) int {
	count := 0
	for y, line := range letters {
		for x, char := range line {
			if char == X {
				count += startingFrom(letters, x, y)
			}
		}
	}

	return count
}

// All 8 directions
var offsets = []struct {
	x int
	y int
}{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}

func startingFrom(letters [][]XMAS, x, y int) int {
	count := 0
	for _, dir := range offsets {
		// Can start from 1 since we already checked letters[y][x] == X
		chars := 1
		for i := 1; i < 4; i++ {
			x := x + dir.x*i
			y := y + dir.y*i
			if x < 0 || y < 0 || x >= len(letters[0]) || y >= len(letters) {
				break
			}
			if letters[y][x] == XMAS(i) {
				chars++
			}
		}
		if chars == 4 {
			count++
		}
	}

	return count
}

func solvePart2(letters [][]XMAS) int {
	count := 0
	for y, line := range letters {
		for x, char := range line {
			if char == A && xAround(letters, x, y) {
				count++
			}
		}
	}

	return count
}

func xAround(letters [][]XMAS, x, y int) bool {
	if x == 0 || y == 0 || x >= len(letters[0])-1 || y >= len(letters)-1 {
		return false
	}

	// The 4 cross directions
	tl := letters[y-1][x-1]
	tr := letters[y-1][x+1]
	bl := letters[y+1][x-1]
	br := letters[y+1][x+1]

	// Check tl and br form mas
	if (tl == M && br == S) || (tl == S && br == M) {
		// Check tr and bl form mas
		if (tr == M && bl == S) || (tr == S && bl == M) {
			return true
		}
	}

	return false
}

func parseInput(path string) [][]XMAS {
	file := commons.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	letters := make([][]XMAS, 0)

	for scanner.Scan() {
		line := make([]XMAS, 0)
		for _, c := range scanner.Text() {
			switch c {
			case 'X':
				line = append(line, X)
			case 'M':
				line = append(line, M)
			case 'A':
				line = append(line, A)
			case 'S':
				line = append(line, S)
			}
		}
		letters = append(letters, line)
	}

	return letters
}
