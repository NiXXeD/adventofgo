package day5

import (
	"adventofgo/src/util"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day5() (string, string) {
	data := util.LoadFile("src/day5/day5.txt")
	lines := strings.Split(data, "\n")

	middle := slices.Index(lines, "")
	firstHalf := lines[0:middle]
	secondHalf := lines[middle+1:]

	rules := make([][]string, len(firstHalf))
	for k, rule := range firstHalf {
		rules[k] = strings.Split(rule, "|")
	}

	updates := make([][]string, len(secondHalf))
	for k, rule := range secondHalf {
		updates[k] = strings.Split(rule, ",")
	}

	part1 := 0
	part2 := 0
	for _, update := range updates {
		if isUpdateValid(rules, update) {
			part1 += findMiddle(update)
		} else {
			fixedUpdate := fixUpdate(rules, update)
			part2 += findMiddle(fixedUpdate)
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func isUpdateValid(rules [][]string, update []string) bool {
	for _, rule := range rules {
		x := slices.Index(update, rule[0])
		y := slices.Index(update, rule[1])
		if x >= 0 && y >= 0 && x > y {
			return false
		}
	}
	return true
}

func findMiddle(update []string) int {
	middleIndex := math.Ceil(float64(len(update) / 2))
	middleValue, _ := strconv.Atoi(update[int(middleIndex)])
	return middleValue
}

func fixUpdate(rules [][]string, update []string) []string {
	slice := make([]string, len(update))
	copy(slice, update)

	for _, rule := range rules {
		x := slices.Index(slice, rule[0])
		y := slices.Index(slice, rule[1])
		if x >= 0 && y >= 0 && x > y {
			if x == 0 {
				slice = slice[1:]
			} else if x == len(slice) {
				slice = slice[0 : x-1]
			} else {
				slice = append(slice[:x], slice[x+1:]...)
			}
			slice = append(slice[:y], append([]string{rule[0]}, slice[y:]...)...)
		}
	}

	if !isUpdateValid(rules, slice) {
		slice = fixUpdate(rules, slice)
	}

	return slice
}
