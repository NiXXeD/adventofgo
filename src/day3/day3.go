package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	data, err := os.ReadFile("./src/day3/day3.txt")
	check(err)

	part1 := process(string(data))
	fmt.Println("Part1:", part1)

	cleaned := cleanPart2(string(data))
	part2 := process(cleaned)

	fmt.Println("Part2:", part2)
}

func cleanPart2(data string) string {
	input := data
	output := ""
	for {
		i := strings.Index(input, "don't()")
		if i >= 0 {
			j := strings.Index(input[i:], "do()")
			j += i

			output += input[:i]
			input = input[j+4:]
		} else {
			output += input
			break
		}
	}
	return output
}

func process(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	result := 0
	for _, m := range matches {
		v1, _ := strconv.Atoi(m[1])
		v2, _ := strconv.Atoi(m[2])
		result += v1 * v2
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}