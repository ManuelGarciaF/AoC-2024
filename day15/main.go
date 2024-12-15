package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	layout, dirs := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(layout, dirs))
	fmt.Println("Part 2: ", solvePart2(layout, dirs))
}

func solvePart1(inputLayout [][]byte, dirs []c.Direction) int {
	// Copy slice to avoid modification
	layout := make([][]byte, len(inputLayout))
	var robotPos c.Coord
	for y, row := range inputLayout {
		layout[y] = slices.Clone(row)
		if x := slices.Index(row, '@'); x >= 0 {
			robotPos = c.Coord{X: x, Y: y}
			layout[y][x] = '.' // Remove the marker from the map
		}
	}

	for _, dir := range dirs {
		targetPos := robotPos.Move(dir)
		switch c.IndexMap(layout, targetPos) {
		case '#':
			continue
		case '.':
			robotPos = targetPos
		case 'O':
			if tryPush(&layout, targetPos, dir) {
				robotPos = targetPos
			}
		}
	}

	fmt.Println()
	for y, row := range layout {
		for x := range row {
			fmt.Print(string(layout[y][x]))
		}
		fmt.Println()
	}

	sum := 0
	for y, row := range layout {
		for x, ch := range row {
			if ch == 'O' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func tryPush(coords *[][]byte, pos c.Coord, dir c.Direction) bool {
	target := pos.Move(dir)
	targetObj := c.IndexMap(*coords, target)
	if targetObj == '.' || (targetObj == 'O' && tryPush(coords, target, dir)) {
		c.SetMap(coords, target, 'O')
		c.SetMap(coords, pos, '.')
		return true
	}

	return false
}

func solvePart2(inputLayout [][]byte, dirs []c.Direction) int {
	// Create a new slice
	layout := make([][]byte, len(inputLayout))
	var robotPos c.Coord
	for y, row := range inputLayout {
		newRow := make([]byte, 2*len(row))
		for x, ch := range row {
			switch ch {
			case '#':
				newRow[2*x] = '#'
				newRow[2*x+1] = '#'
			case '.':
				newRow[2*x] = '.'
				newRow[2*x+1] = '.'
			case 'O':
				newRow[2*x] = '['
				newRow[2*x+1] = ']'
			case '@':
				robotPos = c.Coord{X: 2 * x, Y: y}
				newRow[2*x] = '.'
				newRow[2*x+1] = '.'
			}
		}
		layout[y] = newRow
	}

	for _, dir := range dirs {
		targetPos := robotPos.Move(dir)
		switch c.IndexMap(layout, targetPos) {
		case '#':
			continue
		case '.':
			robotPos = targetPos
		case '[', ']':
			if canPushWide(&layout, targetPos, dir) {
				pushWide(&layout, targetPos, dir)
				robotPos = targetPos
			}
		}
	}

	fmt.Println()
	for y, row := range layout {
		for x := range row {
			fmt.Print(string(layout[y][x]))
		}
		fmt.Println()
	}

	sum := 0
	for y, row := range layout {
		for x, ch := range row {
			if ch == '[' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func canPushWide(layout *[][]byte, pos c.Coord, dir c.Direction) bool {
	boxLeft, boxRight := findBox(*layout, pos)

	canPushTo := func(target c.Coord) bool {
		obj := c.IndexMap(*layout, target)
		return obj == '.' || ((obj == '[' || obj == ']') && canPushWide(layout, target, dir))
	}

	switch dir {
	case c.LEFT:
		return canPushTo(boxLeft.Move(c.LEFT))
	case c.RIGHT:
		return canPushTo(boxRight.Move(c.RIGHT))
	case c.UP, c.DOWN:
		return canPushTo(boxLeft.Move(dir)) && canPushTo(boxRight.Move(dir))
	}
	panic("unreachable")
}

func findBox(layout [][]byte, pos c.Coord) (c.Coord, c.Coord) {
	var boxLeft c.Coord
	var boxRight c.Coord
	if c.IndexMap(layout, pos) == '[' {
		boxLeft = pos
		boxRight = pos.Move(c.RIGHT)
	} else {
		boxLeft = pos.Move(c.LEFT)
		boxRight = pos
	}
	return boxLeft, boxRight
}

func pushWide(target *[][]byte, pos c.Coord, dir c.Direction) {
	obj := c.IndexMap(*target, pos)
	if obj != '[' && obj != ']' {
		return
	}

	boxLeft, boxRight := findBox(*target, pos)

	switch dir {
	case c.LEFT:
		pushWide(target, boxLeft.Move(c.LEFT), c.LEFT)
		c.SetMap(target, boxRight, '.')
		c.SetMap(target, boxLeft, ']')
		c.SetMap(target, boxLeft.Move(c.LEFT), '[')
	case c.RIGHT:
		pushWide(target, boxRight.Move(c.RIGHT), c.RIGHT)
		c.SetMap(target, boxLeft, '.')
		c.SetMap(target, boxRight, '[')
		c.SetMap(target, boxRight.Move(c.RIGHT), ']')
	case c.UP, c.DOWN:
		pushWide(target, boxLeft.Move(dir), dir)
		pushWide(target, boxRight.Move(dir), dir)
		c.SetMap(target, boxLeft, '.')
		c.SetMap(target, boxRight, '.')
		c.SetMap(target, boxLeft.Move(dir), '[')
		c.SetMap(target, boxRight.Move(dir), ']')
	}
}

func parseInput(path string) ([][]byte, []c.Direction) {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	layout := make([][]byte, 0)
	for y := 0; scanner.Scan(); y++ {
		if scanner.Text() == "" {
			break
		}

		layout = append(layout, slices.Clone(scanner.Bytes()))
	}

	directions := make([]c.Direction, 0)
	for scanner.Scan() {
		for _, ch := range scanner.Bytes() {
			switch ch {
			case '^':
				directions = append(directions, c.UP)
			case 'v':
				directions = append(directions, c.DOWN)
			case '<':
				directions = append(directions, c.LEFT)
			case '>':
				directions = append(directions, c.RIGHT)
			}
		}
	}

	return layout, directions
}
