package day2

import (
	"adventofgo/src/util"
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	data := util.LoadFile("./src/day2/day2.txt")
	reports := strings.Split(string(data), "\n")

	var part1 int
	var part2 int
	for _, strLevel := range reports {
		splitLevel := strings.Split(strLevel, " ")
		level := strArrayToIntArray(splitLevel)
		if validateLevelPart1(level) {
			part1++
			part2++
		} else if validateLevelPart2(level) {
			part2++
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func validateLevelPart1(level []int) bool {
	prev := 0
	for k := range level {
		if k > 0 {
			value := level[k] - level[k-1]
			if prev == 0 && ((value > 0 && value < 4) || (value < 0 && value > -4)) {
				// valid
			} else if prev > 0 && value > 0 && value < 4 {
				// valid
			} else if prev < 0 && value < 0 && value > -4 {
				// valid
			} else {
				return false
			}
			prev = value
		}
	}
	return true
}

func validateLevelPart2(level []int) bool {
	for k := range level {
		slice := make([]int, len(level))
		copy(slice, level)

		if k == 0 {
			slice = slice[1:]
		} else if k == len(level) {
			slice = slice[0 : k-1]
		} else {
			slice = append(slice[:k], slice[k+1:]...)
		}
		if validateLevelPart1(slice) {
			return true
		}
	}
	return false
}

func strArrayToIntArray(strArray []string) []int {
	result := make([]int, len(strArray))
	for k, v := range strArray {
		result[k], _ = strconv.Atoi(v)
	}
	return result
}
