package main

import (
	"testing"
)

func BenchmarkSolvePart2(b *testing.B) {
	obstacles, startingPosition, xSize, ySize := parseInput("./input")

    for i := 0; i < b.N; i++ {
        solvePart2(obstacles, startingPosition, xSize, ySize)
    }
}
