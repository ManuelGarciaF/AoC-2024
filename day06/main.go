package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"sync"
	"sync/atomic"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	obstacles, startingPosition, xSize, ySize := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(obstacles, startingPosition, xSize, ySize))
	fmt.Println("Part 2: ", solvePart2(obstacles, startingPosition, xSize, ySize))
}

func solvePart1(obstacles c.Set[c.Coord], startingPosition c.Coord, xSize, ySize int) int {
	visited := path(obstacles, startingPosition, xSize, ySize)

	return visited.Size()
}

func path(obstacles c.Set[c.Coord], startingPosition c.Coord, xSize int, ySize int) c.Set[c.Coord] {
	visited := c.NewSet[c.Coord]()

	curr := startingPosition
	currDir := c.UP
	for curr.X >= 0 && curr.Y >= 0 && curr.X <= xSize && curr.Y <= ySize {
		for obstacles.Contains(curr.Move(currDir)) {
			currDir = c.RotateClockwise[currDir]
		}
		visited.Add(curr)
		curr = curr.Move(currDir)
	}
	return visited
}

func solvePart2(obstacles c.Set[c.Coord], startingPosition c.Coord, xSize, ySize int) int {
	visited := path(obstacles, startingPosition, xSize, ySize)

	var sum atomic.Int32 // Atomic counter for results
	var wg sync.WaitGroup

	// Try adding an obstacle at every step
	for pos := range visited {
		// Can't add an obstacle on starting position
		if pos.Equals(startingPosition) {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()

			obs2 := maps.Clone(obstacles).Add(pos)
			if loops(obs2, startingPosition, xSize, ySize) {
				sum.Add(1)
			}
		}()
	}

	wg.Wait()
	return int(sum.Load())
}

type Step struct {
	c c.Coord
	d c.Direction
}

func loops(obstacles c.Set[c.Coord], startingPosition c.Coord, xSize int, ySize int) bool {
	visited := c.NewSet[Step]()

	curr := startingPosition
	currDir := c.UP
	for curr.X >= 0 && curr.Y >= 0 && curr.X <= xSize && curr.Y <= ySize {
		for obstacles.Contains(curr.Move(currDir)) {
			currDir = c.RotateClockwise[currDir]
		}
		step := Step{c: curr, d: currDir}

		if visited.Contains(step) {
			return true
		}

		visited.Add(step)
		curr = curr.Move(currDir)
	}
	return false
}

// Obstacles, starting position of the guard, and size of the board
func parseInput(path string) (c.Set[c.Coord], c.Coord, int, int) {
	file := c.Must(os.Open(path))
	defer file.Close()

	var startingPosition c.Coord
	obstacles := c.NewSet[c.Coord]()
	scanner := bufio.NewScanner(file)
	y := 0
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		maxY = max(y, maxY)
		for x, ch := range scanner.Bytes() {
			maxX = max(x, maxX)
			switch ch {
			case '#':
				obstacles.Add(c.Coord{X: x, Y: y})
			case '^':
				startingPosition = c.Coord{X: x, Y: y}
			}
		}
		y++
	}

	return obstacles, startingPosition, maxX, maxY
}
