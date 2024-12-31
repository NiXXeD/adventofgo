package main

import (
	"adventofgo/src/day1"
	"adventofgo/src/day2"
	"adventofgo/src/day3"
	"adventofgo/src/day4"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	fmt.Println("Running Day:", day)

	part1 := ""
	part2 := ""

	switch day {
	case "1":
		part1, part2 = day1.Day1()
	case "2":
		part1, part2 = day2.Day2()
	case "3":
		part1, part2 = day3.Day3()
	case "4":
		part1, part2 = day4.Day4()
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
