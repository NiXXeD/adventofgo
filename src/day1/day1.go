package day1

import (
	"adventofgo/src/util"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1() (string, string) {
	data := util.LoadFile("./src/day1/day1.txt")
	lines := strings.Split(data, "\n")

	var left, right []int
	for _, value := range lines {
		values := strings.Split(value, "   ")
		lval, _ := strconv.Atoi(values[0])
		rval, _ := strconv.Atoi(values[1])
		left = append(left, lval)
		right = append(right, rval)
	}

	sort.Ints(left)
	sort.Ints(right)

	var part1 int
	for i := range left {
		value := int(math.Abs(float64(left[i] - right[i])))
		part1 += value
	}

	var part2 int
	for _, num := range left {
		count := 0
		for _, val := range right {
			if num == val {
				count++
			}
		}
		part2 += count * num
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
