package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	heights := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(heights))
	fmt.Println("Part 2: ", solvePart2(heights))
}

func solvePart1(heights [][]int) int {
	trailheads := make([]c.Coord, 0)
	for y, row := range heights {
		for x, n := range row {
			if n == 0 {
				trailheads = append(trailheads, c.Coord{X: x, Y: y})
			}
		}
	}
	xSize := len(heights[0]) - 1
	ySize := len(heights) - 1

	count := 0
	for _, trailhead := range trailheads {
		count += score(trailhead, heights, xSize, ySize, c.NewSet[c.Coord]()).Size()
	}

	return count
}

func score(position c.Coord, heights [][]int, xSize, ySize int, reached c.Set[c.Coord]) c.Set[c.Coord] {
	val := c.IndexMap(heights, position)

	if val == 9 {
		// Maybe don't even need to clone
		return reached.Clone().Add(position)
	}

	moves := []c.Coord{
		position.Move(c.UP),
		position.Move(c.DOWN),
		position.Move(c.LEFT),
		position.Move(c.RIGHT),
	}

	// FP gets really ugly in go
	return c.Foldl(
		reached,
		moves,
		func(reached c.Set[c.Coord], move c.Coord) c.Set[c.Coord] {
			if !move.Inbounds(xSize, ySize) || c.IndexMap(heights, move) != val+1 {
				return reached
			}
			return score(move, heights, xSize, ySize, reached)
		},
	)
}

func rating(position c.Coord, heights [][]int, xSize, ySize int) int {
	val := c.IndexMap(heights, position)

	if val == 9 {
		return 1
	}

	moves := []c.Coord{
		position.Move(c.UP),
		position.Move(c.DOWN),
		position.Move(c.LEFT),
		position.Move(c.RIGHT),
	}

	return c.Sum(c.Map(
		moves,
		func(move c.Coord) int {
			if !move.Inbounds(xSize, ySize) || c.IndexMap(heights, move) != val+1 {
				return 0
			}
			return rating(move, heights, xSize, ySize)
		},
	))
}

func solvePart2(heights [][]int) int {
	trailheads := make([]c.Coord, 0)
	for y, row := range heights {
		for x, n := range row {
			if n == 0 {
				trailheads = append(trailheads, c.Coord{X: x, Y: y})
			}
		}
	}
	xSize := len(heights[0]) - 1
	ySize := len(heights) - 1

	count := 0
	for _, trailhead := range trailheads {
		count += rating(trailhead, heights, xSize, ySize)
	}

	return count
}

func parseInput(path string) [][]int {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heights := make([][]int, 0)

	y := 0
	for scanner.Scan() {
		heights = append(heights,
			c.Map(
				slices.Clone(scanner.Bytes()),
				func(ch byte) int {
					return c.MustAtoi(string(ch))
				},
			),
		)
		y++
	}

	return heights
}
