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
	switch day {
	case "1":
		day1.Day1()
	case "2":
		day2.Day2()
	case "3":
		day3.Day3()
	case "4":
		day4.Day4()
	}
}
