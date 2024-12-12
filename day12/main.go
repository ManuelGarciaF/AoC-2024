package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	plots := parseInput(os.Args[1])
	regions := groupRegions(plots)

	fmt.Println("Part 1: ", solvePart1(regions))
	fmt.Println("Part 2: ", solvePart2(regions))
}

func groupRegions(plots [][]byte) []c.Set[c.Coord] {
	// Separate crops
	crops := make(map[byte][]c.Coord)
	for y, row := range plots {
		for x, b := range row {
			crops[b] = append(crops[b], c.Coord{X: x, Y: y})
		}
	}

	// Separate regions
	regions := make([]c.Set[c.Coord], 0)
	for _, crop := range crops {
		regions = append(regions, cropRegions(crop)...)
	}
	return regions
}

func cropRegions(coords []c.Coord) []c.Set[c.Coord] {
	regions := make([]c.Set[c.Coord], 0)
	for _, coord := range coords {
		matched := false
		matchedIndex := -1
		for i := 0; i < len(regions); i++ {
			if edgesShared(coord, regions[i]) > 0 {
				if matched {
					// Found a different region that is also in contact with the current coord,
					// they are the same region, need to merge.
					regions[matchedIndex].Union(regions[i])
					regions = slices.Delete(regions, i, i+1)
					i--
				} else {
					regions[i].Add(coord)
					matched = true
					matchedIndex = i
				}
			}
		}
		if !matched {
			regions = append(regions, c.NewSet[c.Coord]().Add(coord))
		}
	}

	return regions
}

func edgesShared(coord c.Coord, area c.Set[c.Coord]) int {
	candidates := []c.Coord{
		coord.Move(c.UP),
		coord.Move(c.DOWN),
		coord.Move(c.LEFT),
		coord.Move(c.RIGHT),
	}
	return len(c.Filter(candidates, area.Contains))
}

func solvePart1(regions []c.Set[c.Coord]) int {
	cost := 0
	for _, region := range regions {
		perimeter := 0
		for coord := range region {
			perimeter += 4 - edgesShared(coord, region)
		}

		cost += perimeter * region.Size()
	}

	return cost
}

type Edge struct {
	c.Coord
	c.Direction
}

func solvePart2(regions []c.Set[c.Coord]) int {
	cost := 0
	for _, region := range regions {
		cost += numberOfEdges(region) * region.Size()
	}

	return cost
}

func numberOfEdges(region c.Set[c.Coord]) int {
	// Find individual edges of every point
	individualEdges := c.NewSet[Edge]()
	for coord := range region {
		for _, d := range c.Directions {
			if region.Contains(coord.Move(d)) {
				continue
			}
			individualEdges.Add(Edge{coord, d})
		}
	}

	// Join edges that are contiguous into groups
	edgeGroups := make([]c.Set[Edge], 0)
	for ie := range individualEdges {
		matched := false
		matchedIndex := -1
		for i := 0; i < len(edgeGroups); i++ {
			// The possible contiguous edges
			orthDirs := c.OrthogonalDirections[ie.Direction]
			orthCoord1 := ie.Move(orthDirs[0])
			orthCoord2 := ie.Move(orthDirs[1])
			if edgeGroups[i].Contains(Edge{orthCoord1, ie.Direction}) || edgeGroups[i].Contains(Edge{orthCoord2, ie.Direction}) {
				// Need to merge two edges into one
				if matched {
					edgeGroups[matchedIndex].Union(edgeGroups[i])
					edgeGroups = slices.Delete(edgeGroups, i, i+1)
					i--
				}
				edgeGroups[i].Add(ie)
				matched = true
				matchedIndex = i
			}
		}
		// If it wasn't part of any previous edge, add a new one
		if !matched {
			edgeGroups = append(edgeGroups, c.NewSet[Edge]().Add(ie))
		}
	}

	return len(edgeGroups)
}

func parseInput(path string) [][]byte {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plots := make([][]byte, 0)
	for scanner.Scan() {
		plots = append(plots,
			slices.Clone(scanner.Bytes()),
		)
	}

	return plots
}
