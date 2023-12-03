package main

import (
	"strconv"
	"strings"
)

var colorLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// I'm using the world's most inefficient parser for this question :P. Hand rolling isn't that hard.. should prob just do that.
// Part 1 and part 2 still run within hundreds of micros on my M2, though, so I guess it's fine. I don't care enough to make it go faster.
//
// Yes, I'm also not handling errors because Go is infuriatingly verbose. They should never happen if our input is correct (we control it, so it is...).
type Day2 struct {
	filePath string
}

func (Day2) Name() string {
	return "day2"
}
func (day Day2) FilePath() string {
	return day.filePath
}
func (Day2) Part1(lines []string) (int, error) {
	colorFrequency := make(map[string]int, 3)
	total := 0
	for i, line := range lines {
		// assume they're in order since input is actually in order. get to skip a parsing step.
		gameId := i + 1
		_, game, _ := strings.Cut(line, ": ")
		rounds := strings.Split(game, "; ")
		limitReached := false
		for _, round := range rounds {
			colorFreqPairs := strings.Split(round, ", ")
			for _, pair := range colorFreqPairs {
				freq, color, _ := strings.Cut(pair, " ")
				num, _ := strconv.Atoi(freq)
				if existing, ok := colorFrequency[color]; ok {
					colorFrequency[color] = existing + num
				} else {
					colorFrequency[color] = num
				}
			}

			// evaluate round
			for color, limit := range colorLimits {
				existing, ok := colorFrequency[color]
				if ok && existing > limit {
					// fmt.Println("Game", gameId, "adding color", color, "due to limit being", limit)
					limitReached = true
					break
				}
			}
			clear(colorFrequency)
			if limitReached {
				break
			}
		}

		if !limitReached {
			total += gameId
			continue
		}
	}
	return total, nil
}
func (Day2) Part2(lines []string) (int, error) {
	colorFrequency := make(map[string]int, 3)
	maxFrequency := make(map[string]int, 3)
	total := 0
	for _, line := range lines {
		// assume they're in order since input is actually in order. get to skip a parsing step.
		_, game, _ := strings.Cut(line, ": ")
		rounds := strings.Split(game, "; ")
		for _, round := range rounds {
			colorFreqPairs := strings.Split(round, ", ")
			for _, pair := range colorFreqPairs {
				freq, color, _ := strings.Cut(pair, " ")
				num, _ := strconv.Atoi(freq)
				if existing, ok := colorFrequency[color]; ok {
					colorFrequency[color] = existing + num
				} else {
					colorFrequency[color] = num
				}
			}

			// evaluate round
			for color, freq := range colorFrequency {
				existing, ok := maxFrequency[color]
				if !ok {
					maxFrequency[color] = freq
				} else if freq > existing {
					maxFrequency[color] = freq
				}
			}
			clear(colorFrequency)
		}

		power := 1
		for _, freq := range maxFrequency {
			power *= freq
		}
		total += power

		// must clear after every game
		clear(maxFrequency)
	}
	return total, nil
}
