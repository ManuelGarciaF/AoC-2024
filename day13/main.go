package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

type Machine struct {
	Ax, Ay, Bx, By int
	PrizeX, PrizeY int
}

func main() {
	machines := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(machines))
	fmt.Println("Part 2: ", solvePart2(machines))
}

func solvePart1(machines []Machine) int {
	tokens := 0
	for _, machine := range machines {
		amount, possible := tokensFor(machine)
		if possible {
			tokens += amount
		}
	}
	return tokens
}

func solvePart2(machines []Machine) int {
	tokens := 0
	for _, machine := range machines {
		machine.PrizeX += 10000000000000
		machine.PrizeY += 10000000000000
		amount, possible := tokensFor(machine)
		if possible {
			tokens += amount
		}
	}
	return tokens
}

func tokensFor(m Machine) (int, bool) {
	x := float64(m.PrizeX)
	y := float64(m.PrizeY)
	aX := float64(m.Ax)
	aY := float64(m.Ay)
	bX := float64(m.Bx)
	bY := float64(m.By)

	div := aX*bY-aY*bX
	a := (x*bY-y*bX)/div
	b := (y*aX-x*aY)/div

	// I don't know why div is always != 0, isn't there supposed to be more than one way to
	// get to the prize? Why does it mention "the fewest number of moves" if there is
	// always just one solution.

	// Only integer solutions are valid
	if a != math.Floor(a) || b != math.Floor(b) {
		return -1, false
	}

	return int(a*3+b), true
}

func parseInput(path string) []Machine {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Modified SplitLines function from the bufio source to split by paragraphs
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := strings.Index(string(data), "\n\n"); i >= 0 {
			return i + 2, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	})

	machines := make([]Machine, 0)

	re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

	for scanner.Scan() {
		results := re.FindStringSubmatch(scanner.Text())
		machines = append(machines, Machine{
			Ax:     c.MustAtoi(results[1]),
			Ay:     c.MustAtoi(results[2]),
			Bx:     c.MustAtoi(results[3]),
			By:     c.MustAtoi(results[4]),
			PrizeX: c.MustAtoi(results[5]),
			PrizeY: c.MustAtoi(results[6]),
		})
	}

	return machines
}
