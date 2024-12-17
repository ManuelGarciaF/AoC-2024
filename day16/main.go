package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	walls, start, end := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(walls, start, end))
	fmt.Println("Part 2: ", solvePart2(walls, start, end))
}

type Node struct {
	c.Coord
	c.Direction
}

type Result struct {
	dist  int
	prevs []c.Coord
}

func djakstra(walls [][]bool, start c.Coord) map[c.Coord]Result {
	pq := c.NewPriorityQueue[Node]()
	pq.PushItem(Node{Coord: start, Direction: c.RIGHT}, 0)

	results := make(map[c.Coord]Result)
	for !pq.IsEmpty() {
		node, dist := pq.PopItem()

		// Skip if the current distance is worse than the best known
		if result, visited := results[node.Coord]; visited && dist > result.dist {
			continue
		}

		neighbors := []struct {
			dir  c.Direction
			cost int
		}{
			{node.Direction, 1},                       // Forwards
			{c.RotateClockwise[node.Direction], 1001}, // Rotations
			{c.RotateCounterClockwise[node.Direction], 1001},
		}

		for _, n := range neighbors {
			next := node.Move(n.dir)
			if c.IndexMap(walls, next) { // There is a wall, invalid
				continue
			}

			nextDist := dist + n.cost
			result, visited := results[next]

			if visited && result.dist == nextDist {
				// Add as another predecesor
				result.prevs = append(result.prevs, node.Coord)
				results[next] = result
				continue
			}

			if visited && result.dist < nextDist { // There is already a shorter path
				continue
			}

			// It wasn't visited or found a better path.
			results[next] = Result{dist: nextDist, prevs: []c.Coord{node.Coord}}
			pq.PushItem(Node{Coord: next, Direction: n.dir}, nextDist)
		}
	}

	return results
}

func solvePart1(walls [][]bool, start, end c.Coord) int {
	return djakstra(walls, start)[end].dist
}

func solvePart2(walls [][]bool, start, end c.Coord) int {
	results := djakstra(walls, start)

	tiles := c.NewSet[c.Coord]()
	q := make(c.Queue[c.Coord], 0)
	q.Push(end)
	for !q.IsEmpty() {
		c := q.Pop()
		tiles.Add(c)
		q.Push(results[c].prevs...)
	}

	for y, row := range walls {
		for x, wall := range row {
			if wall {
				fmt.Print("███")
			} else {
				if tiles.Contains(c.Coord{X: x, Y: y}) {
					fmt.Printf("✔%2d", (results[c.Coord{X: x, Y: y}].dist/1000)%100)
				} else {
					fmt.Printf(" %2d", (results[c.Coord{X: x, Y: y}].dist/1000)%100)
				}
			}
		}
		fmt.Println()
	}

	return tiles.Size()
}

func parseInput(path string) ([][]bool, c.Coord, c.Coord) {
	file := c.Must(os.Open(path))
	defer file.Close()
	scanner := bufio.NewScanner(file)

	walls := make([][]bool, 0)
	var start, end c.Coord
	for y := 0; scanner.Scan(); y++ {
		walls = append(walls, c.Map(scanner.Bytes(), func(ch byte) bool {
			return ch == '#'
		}))
		if x := slices.Index(scanner.Bytes(), 'S'); x >= 0 {
			start = c.Coord{X: x, Y: y}
		}
		if x := slices.Index(scanner.Bytes(), 'E'); x >= 0 {
			end = c.Coord{X: x, Y: y}
		}
	}

	return walls, start, end
}
