package adventofgo

import (
	"adventofgo/src/day1"
	"adventofgo/src/day2"
	"adventofgo/src/day3"
	"adventofgo/src/day4"
	"testing"
)

type DayFn func() (string, string)

func TestDays(t *testing.T) {
	tests := []struct {
		fn    DayFn
		day   string
		part1 string
		part2 string
	}{
		{day1.Day1, "1", "1666427", "24316233"},
		{day2.Day2, "2", "299", "364"},
		{day3.Day3, "3", "178538786", "102467299"},
		{day4.Day4, "4", "2573", "1850"},
	}

	for _, test := range tests {
		t.Run("Days", func(t *testing.T) {
			result1, result2 := test.fn()
			if result1 != test.part1 {
				t.Errorf("Day %v Part %v test failed, expected %v, got %v", test.day, 1, test.part1, result1)
			}
			if result2 != test.part2 {
				t.Errorf("Day %v Part %v test failed, expected %v, got %v", test.day, 2, test.part2, result2)
			}
		})
	}
}
