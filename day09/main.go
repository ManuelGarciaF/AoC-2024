package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	c "github.com/ManuelGarciaF/AoC-2024/commons"
)

func main() {
	content := parseInput(os.Args[1])

	fmt.Println("Part 1: ", solvePart1(content))
	fmt.Println("Part 2: ", solvePart2(content))
}

func solvePart1(diskMap []byte) uint64 {
	layout := expand(diskMap)

	// Compact disk
	end := len(layout) - 1
	for i, ch := range layout {
		if i >= end { // Stop after reaching the new end
			break
		}

		// Ignore if not empty
		if ch != -1 {
			continue

		}

		id := -1
		for id == -1 { // Go back as far as needed
			id = layout[end]
			end--
		}

		// Update place
		layout[i] = id
	}

	layout = layout[:end]

	// Compute checksum (could be done during prev step)
	sum := uint64(0)
	for i, ch := range layout {
		sum += uint64(i) * uint64(ch)
	}

	return sum
}

// Returns what id is at each byte, using -1 for empty
func expand(diskMap []byte) []int {
	expanded := make([]int, 0, 10*len(diskMap))

	add := func(ch int, count byte) {
		for i := byte(0); i < count; i++ {
			expanded = append(expanded, ch)
		}
	}

	id := 0
	for i, count := range diskMap {
		if i%2 == 0 { // even -> file
			add(id, count)
			id++
		} else { // odd -> empty
			add(-1, count)
		}
	}
	return expanded
}

type Partition struct {
	Size int
	Id   int // -1 for free partitions
}

func solvePart2(diskMap []byte) uint64 {
	partitions := partitionList(diskMap)

	for i := len(partitions) - 1; i > 0; i-- {
		fileP := partitions[i]
		if fileP.Id == -1 { // Don't do anything with free partitions
			continue
		}

		// Check if there is free space at the start
		for j := 0; j < i; j++ {
			freeP := partitions[j]
			if freeP.Id != -1 {
				continue
			}
			if freeP.Size < fileP.Size {
				continue
			}

			// Found a free partition with enough space
			if freeP.Size == fileP.Size { // No remaining free space
				partitions[j].Id = fileP.Id
			} else { // Have to divide the free partition
				remaining := freeP.Size - fileP.Size
				partitions = slices.Insert(partitions, j, fileP) // Insert the file before the partition
				partitions[j+1].Size = remaining
				// Fix counters since we just added an element
				i++
			}

			// Mark the original partition as free
			partitions[i].Id = -1

			break // We already found a valid free partition
		}
	}

	i := uint64(0)
	sum := uint64(0)
	for _, p := range partitions {
		for j := 0; j < p.Size; j++ {
			if p.Id > -1 {
				sum += i * uint64(p.Id)
			}
			i++
		}
	}

	return sum
}

func partitionList(diskMap []byte) []Partition {
	partitions := make([]Partition, 0, len(diskMap))

	id := 0
	for i, size := range diskMap {
		if i%2 == 0 { // even -> file
			partitions = append(partitions, Partition{
				Size: int(size),
				Id:   id,
			})
			id++
		} else { // odd -> empty
			partitions = append(partitions, Partition{
				Size: int(size),
				Id:   -1,
			})
		}
	}

	return partitions
}

func printPartitions(ps []Partition) {
	for _, p := range ps {
		for i := 0; i < p.Size; i++ {
			if p.Id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(p.Id)
			}
		}
	}
	fmt.Println()
}

func parseInput(path string) []byte {
	file := c.Must(os.Open(path))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	chars := scanner.Bytes()
	diskMap := make([]byte, 0, len(chars))
	for _, ch := range chars {
		diskMap = append(diskMap, byte(c.MustAtoi(string(ch))))
	}

	return diskMap
}
