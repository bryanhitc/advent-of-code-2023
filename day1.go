package main

import (
	"errors"
	"strings"
	"unicode"
)

var digitWordMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var notFoundErr = errors.New("404: Not found")

type Day1 struct {
	filePath string
}

func (Day1) Name() string {
	return "day1"
}
func (day Day1) FilePath() string {
	return day.filePath
}
func (Day1) Part1(lines []string) (int, error) {
	calculateLine := func(line string) int {
		first_digit := 0
		last_digit := 0
		found_first := false
		for _, char := range line {
			if !unicode.IsDigit(char) {
				continue
			}
			digit := int(char - '0')
			if !found_first {
				first_digit = digit
			}
			last_digit = digit
			found_first = true
		}
		return (first_digit * 10) + last_digit
	}

	total := 0
	for _, line := range lines {
		total += calculateLine(line)
	}
	return total, nil
}
func (Day1) Part2(lines []string) (int, error) {
	extractDigit := func(line string, i int, char rune) (int, error) {
		if unicode.IsDigit(char) {
			return int(char - '0'), nil
		}

		// Inefficient. I don't care.
		for word, val := range digitWordMap {
			if strings.HasPrefix(line[i:], word) {
				return val, nil
			}
		}
		return 0, notFoundErr
	}
	calculateLine := func(line string) int {
		first_digit := 0
		last_digit := 0
		found_first := false
		for i, char := range line {
			digit, err := extractDigit(line, i, char)
			if err != nil {
				continue
			}
			if !found_first {
				first_digit = digit
			}
			last_digit = digit
			found_first = true
		}
		return (first_digit * 10) + last_digit
	}

	total := 0
	for _, line := range lines {
		total += calculateLine(line)
	}
	return total, nil
}
