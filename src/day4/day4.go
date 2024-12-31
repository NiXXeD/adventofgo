package day4

import (
	"adventofgo/src/util"
	"strconv"
	"strings"
)

func Day4() (string, string) {
	data := util.LoadFile("./src/day4/day4.txt")
	lines := strings.Split(data, "\n")

	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		splitline := strings.Split(line, "")
		for x, cell := range splitline {
			grid[y][x] = cell
		}
	}

	part1 := 0
	part2 := 0
	for y := range len(grid) {
		for x := range grid[y] {
			part1 += findXmas(grid, x, y)
			part2 += findCrossmas(grid, x, y)
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func findXmas(grid [][]string, x int, y int) int {
	count := 0

	// Right
	count += findWordInDirection(grid, x, y, 1, 0)
	// Right Down
	count += findWordInDirection(grid, x, y, 1, 1)
	// Down
	count += findWordInDirection(grid, x, y, 0, 1)
	// Down Left
	count += findWordInDirection(grid, x, y, -1, 1)
	// Left
	count += findWordInDirection(grid, x, y, -1, 0)
	// Left Up
	count += findWordInDirection(grid, x, y, -1, -1)
	// Up
	count += findWordInDirection(grid, x, y, 0, -1)
	// Up Right
	count += findWordInDirection(grid, x, y, 1, -1)

	return count
}

func findWordInDirection(grid [][]string, x int, y int, xd int, yd int) int {
	if safeGrid(grid, x, y) == "X" {
		if safeGrid(grid, x+xd, y+yd) == "M" {
			if safeGrid(grid, x+(xd*2), y+(yd*2)) == "A" {
				if safeGrid(grid, x+(xd*3), y+(yd*3)) == "S" {
					return 1
				}
			}
		}
	}
	return 0
}

func findCrossmas(grid [][]string, x int, y int) int {
	count := 0

	// Right
	count += checkCross(grid, x, y, "M", "S", "M", "S")
	// Down
	count += checkCross(grid, x, y, "M", "M", "S", "S")
	// Left
	count += checkCross(grid, x, y, "S", "M", "S", "M")
	// Up
	count += checkCross(grid, x, y, "S", "S", "M", "M")

	return count
}

func checkCross(grid [][]string, x int, y int, ul string, ur string, bl string, br string) int {
	if safeGrid(grid, x, y) == "A" {
		if safeGrid(grid, x-1, y-1) == ul {
			if safeGrid(grid, x-1, y+1) == bl {
				if safeGrid(grid, x+1, y-1) == ur {
					if safeGrid(grid, x+1, y+1) == br {
						return 1
					}
				}
			}
		}
	}

	return 0
}

func safeGrid(grid [][]string, x int, y int) string {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		return grid[y][x]
	}
	return ""
}
