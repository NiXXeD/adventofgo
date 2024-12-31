package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day1.txt")
	check(err)

	lines := strings.Split(string(data), "\n")

	var left, right []int
	for _, value := range lines {
		values := strings.Split(value, "   ")
		lval, err := strconv.Atoi(values[0])
		check(err)
		rval, err := strconv.Atoi(values[1])
		check(err)
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

	fmt.Println("Part1:", part1)

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

	fmt.Println("Part2:", part2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
