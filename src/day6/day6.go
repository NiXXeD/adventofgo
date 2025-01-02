package day6

import (
	"adventofgo/src/util"
	"fmt"
	"strconv"
	"strings"
)

func Day6() (string, string) {
	data := util.LoadFile("src/day6/day6.txt")
	lines := strings.Split(data, "\n")

	var gx, gy int
	// Up
	gxd := 0
	gyd := -1

	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		splitline := strings.Split(line, "")
		for x, cell := range splitline {
			if cell == "^" {
				gx = x
				gy = y
				cell = "X"
			}
			grid[y][x] = cell
		}
	}

	part1 := 1
	for {
		nx := gx + gxd
		ny := gy + gyd
		val := safeGrid(grid, nx, ny)
		fmt.Println(nx, ny, val)
		if val == "." {
			part1++
			gx = nx
			gy = ny
			grid[ny][nx] = "X"
		} else if val == "X" {
			gx = nx
			gy = ny
		} else if val == "#" {
			if gxd == 0 && gyd == -1 {
				// Right
				gxd = 1
				gyd = 0
			} else if gxd == 1 && gyd == 0 {
				// Down
				gxd = 0
				gyd = 1
			} else if gxd == 0 && gyd == 1 {
				// Left
				gxd = -1
				gyd = 0
			} else if gxd == -1 && gyd == 0 {
				// Up
				gxd = 0
				gyd = -1
			}
		} else {
			// Exit
			break
		}
	}

	part2 := 0

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func safeGrid(grid [][]string, x int, y int) string {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		return grid[y][x]
	}
	return ""
}
