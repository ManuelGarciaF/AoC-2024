package main

import (
	"bufio"
	"fmt"
	"os"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

type Antenna struct {
	position  c.Coord
	frequency byte
}

func main() {
	antennas, xSize, ySize := parseInput(os.Args[1])

	// Group antennas by frequency
	groups := make(map[byte][]c.Coord)

	for _, a := range antennas {
		groups[a.frequency] = append(groups[a.frequency], a.position)
	}

	fmt.Println("Part 1: ", solvePart1(groups, xSize, ySize))
	fmt.Println("Part 2: ", solvePart2(groups, xSize, ySize))
}

func solvePart1(groups map[byte][]c.Coord, xSize, ySize int) int {
	antinodes := c.NewSet[c.Coord]()
	for _, positions := range groups {
		antinodes.Union(firstAntinodesOf(positions, xSize, ySize))
	}

	return antinodes.Size()
}

func solvePart2(groups map[byte][]c.Coord, xSize, ySize int) int {
	antinodes := c.NewSet[c.Coord]()
	for _, positions := range groups {
		antinodes.Union(allAntinodesOf(positions, xSize, ySize))
	}

	return antinodes.Size()
}

func firstAntinodesOf(positions []c.Coord, xSize, ySize int) c.Set[c.Coord] {
	antinodes := c.NewSet[c.Coord]()
	for _, a1 := range positions {
		for _, a2 := range positions {
			if a1.Equals(a2) {
				continue
			}

			// Just add vectors lol
			ant1 := a2.Scale(2).Sub(a1)
			ant2 := a1.Scale(2).Sub(a2)

			if ant1.Inbounds(xSize, ySize) {
				antinodes.Add(ant1)
			}
			if ant2.Inbounds(xSize, ySize) {
				antinodes.Add(ant2)
			}
		}
	}

	return antinodes
}

func allAntinodesOf(positions []c.Coord, xSize, ySize int) c.Set[c.Coord] {
	antinodes := c.NewSet[c.Coord]()
	for _, a1 := range positions {
		for _, a2 := range positions {
			if a1.Equals(a2) {
				continue
			}

			// Just add vectors lol
			// Go in one direction
			diffVec1 := a2.Sub(a1)
			for i := 1; true; i++ { // If outside of bounds break
				antinode := a1.Add(diffVec1.Scale(i))
				if !antinode.Inbounds(xSize, ySize) {
					break
				}
				antinodes.Add(antinode)
			}

			// Go in the other direction
			diffVec2 := a1.Sub(a2)
			for i := 1; true; i++ { // If outside of bounds break
				antinode := a2.Add(diffVec2.Scale(i))
				if !antinode.Inbounds(xSize, ySize) {
					break
				}
				antinodes.Add(antinode)
			}
		}
	}

	return antinodes
}

func parseInput(path string) ([]Antenna, int, int) {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennas := make([]Antenna, 0)
	y := 0
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		maxY = max(y, maxY)
		for x, ch := range scanner.Bytes() {
			maxX = max(x, maxX)
			if ch == '.' {
				continue
			}

			antennas = append(antennas, Antenna{
				position:  c.Coord{X: x, Y: y},
				frequency: ch,
			})
		}
		y++
	}

	return antennas, maxX, maxY
}
