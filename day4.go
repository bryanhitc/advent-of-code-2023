package main

import (
	"strconv"
	"strings"
)

type Day4 struct {
	filePath string
}

func (Day4) Name() string {
	return "day4"
}
func (day Day4) FilePath() string {
	return day.filePath
}

type Card struct {
	winningNums map[int]Empty
	nums        []int
}

func Parse(lines []string) []Card {
	cards := make([]Card, 0, len(lines))
	for _, line := range lines {
		_, content, _ := strings.Cut(line, ":")
		first, second, _ := strings.Cut(content, "|")

		winningNums := make(map[int]Empty, 0)
		for _, winningNumStr := range strings.Fields(first) {
			num, _ := strconv.Atoi(winningNumStr)
			winningNums[num] = Empty{}
		}

		nums := make([]int, 0)
		for _, numStr := range strings.Fields(second) {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		cards = append(cards, Card{
			winningNums,
			nums,
		})
	}
	return cards
}

func (Day4) Part1(lines []string) (int, error) {
	cards := Parse(lines)
	totalPoints := 0
	for _, card := range cards {
		points := 1
		for _, num := range card.nums {
			if _, ok := card.winningNums[num]; ok {
				points *= 2
			}
		}

		// normalize for points starting at 1 instead of 0. doing a
		// single "bitshift" at the end is much cheaper than introducing
		// an `if points == 0` branch in the for loop.
		points /= 2
		totalPoints += points
	}
	return totalPoints, nil
}

func (Day4) Part2(lines []string) (int, error) {
	cards := Parse(lines)
	cardFreq := make(map[int]int, len(cards))
	for i, card := range cards {
		cardFreq[i] += 1
		numCopies := cardFreq[i]
		numWinning := 0
		for _, num := range card.nums {
			if _, ok := card.winningNums[num]; ok {
				numWinning += 1
				cardFreq[i+numWinning] += numCopies
			}
		}
	}

	totalCards := 0
	for _, freq := range cardFreq {
		totalCards += freq
	}
	return totalCards, nil
}
