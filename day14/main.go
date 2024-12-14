package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"
	"regexp"
	"strings"

	"golang.org/x/image/bmp"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

type Robot struct {
	pos, velocity c.Coord
}

func main() {
	robots := parseInput(os.Args[1])

	xSize, ySize := 101, 103
	if strings.Index(os.Args[1], "sample") >= 0 {
		xSize, ySize = 11, 7
	}

	fmt.Println("Part 1: ", solvePart1(robots, xSize, ySize))

	// WARNING this creates 7000 images
	solvePart2(robots, xSize, ySize)
}

func solvePart1(robots []Robot, xSize, ySize int) int {
	seconds := 100

	halfX := xSize / 2
	halfY := ySize / 2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, r := range robots {
		finalPos := r.pos.Add(r.velocity.Scale(seconds)).WrapAround(xSize, ySize)
		if finalPos.X > halfX && finalPos.Y > halfY {
			q1++
		}
		if finalPos.X < halfX && finalPos.Y > halfY {
			q2++
		}
		if finalPos.X < halfX && finalPos.Y < halfY {
			q3++
		}
		if finalPos.X > halfX && finalPos.Y < halfY {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func solvePart2(robots []Robot, xSize, ySize int) {
	advance := func(r Robot) Robot {
		return Robot{
			pos:      r.pos.Add(r.velocity).WrapAround(xSize, ySize),
			velocity: r.velocity,
		}
	}

	for i := 0; i < 7000; i++ {
		robots = c.Map(robots, advance)
		positions := make(map[c.Coord]int)
		for _, r := range robots {
			positions[r.pos]++
		}

		img := image.NewGray(image.Rect(0, 0, xSize, ySize))

		for y := 0; y < ySize; y++ {
			for x := 0; x < xSize; x++ {
				if positions[c.Coord{X: x, Y: y}] > 0 {
					img.Set(x, y, color.Black)
				} else {
					img.Set(x, y, color.White)
				}
			}
		}

		path := fmt.Sprintf("day14/images/step-%v.bmp", i+1)
		file := c.Must(os.Create(path))
		bmp.Encode(file, img)
		file.Close()
	}
}

func parseInput(path string) []Robot {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)

	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for scanner.Scan() {
		results := re.FindStringSubmatch(scanner.Text())
		robots = append(robots, Robot{
			pos: c.Coord{
				X: c.MustAtoi(results[1]),
				Y: c.MustAtoi(results[2]),
			},
			velocity: c.Coord{
				X: c.MustAtoi(results[3]),
				Y: c.MustAtoi(results[4]),
			},
		})
	}

	return robots
}
