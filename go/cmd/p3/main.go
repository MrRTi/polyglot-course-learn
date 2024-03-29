package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

type SlopeElmt = int

const (
	Tree SlopeElmt = iota
	Snow
)

func main() {
	treeCount := 0
	for r, line := range strings.Split(getInput(), "\n") {
		if string(line[r*3%len(line)]) == "#" {
			treeCount += 1
		}
	}

	fmt.Printf("tree count %v\n", treeCount)
}
