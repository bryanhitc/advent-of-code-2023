package main

import (
	"testing"
)

func runTest(t *testing.T, problem Problem, expectedPart1 int, expectedPart2 int) {
	lines, err := readLines(problem.FilePath())
	if err != nil {
		t.Fatalf("Failed to read input: %s", problem.FilePath())
	}

	part1, err := problem.Part1(lines)
	if err != nil {
		t.Errorf("Part1 unexpected error: %s", err.Error())
	}
	if part1 != expectedPart1 {
		t.Errorf("Part1 returns wrong answer. Expected: %d, actual: %d", expectedPart1, part1)
	}

	part2, err := problem.Part2(lines)
	if err != nil {
		t.Errorf("Part2 unexpected error: %s", err.Error())
	}
	if part2 != expectedPart2 {
		t.Errorf("Part2 returns wrong answer. Expected: %d, actual: %d", expectedPart2, part2)
	}
}

func runBench(b *testing.B, problem Problem) {
	lines, err := readLines(problem.FilePath())
	if err != nil {
		b.Fatalf("Failed to read input: %s", problem.FilePath())
	}
	b.Run("Part1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			problem.Part1(lines)
		}
	})
	b.Run("Part2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			problem.Part2(lines)
		}
	})
}

func TestDay1(t *testing.T) {
	runTest(t, Day1{filePath: "inputs/day1/test1.txt"}, 142, 142)
	runTest(t, Day1{filePath: "inputs/day1/test2.txt"}, 209, 281)
	runTest(t, Day1{filePath: "inputs/day1/problem.txt"}, 54388, 53515)
}

func BenchmarkDay1(b *testing.B) {
	runBench(b, Day1{filePath: "inputs/day1/problem.txt"})
}

func TestDay2(t *testing.T) {
	runTest(t, Day2{filePath: "inputs/day2/test1.txt"}, 8, 2286)
	runTest(t, Day2{filePath: "inputs/day2/problem.txt"}, 2207, 62241)
}

func BenchmarkDay2(b *testing.B) {
	runBench(b, Day2{filePath: "inputs/day2/problem.txt"})
}

func TestDay3(t *testing.T) {
	runTest(t, Day3{filePath: "inputs/day3/test1.txt"}, 4361, 467835)
	runTest(t, Day3{filePath: "inputs/day3/problem.txt"}, 528819, 80403602)
}

func BenchmarkDay3(b *testing.B) {
	runBench(b, Day3{filePath: "inputs/day3/problem.txt"})
}

func TestDay4(t *testing.T) {
	runTest(t, Day4{filePath: "inputs/day4/test1.txt"}, 13, 30)
	runTest(t, Day4{filePath: "inputs/day4/problem.txt"}, 25571, 8805731)
}

func BenchmarkDay4(b *testing.B) {
	runBench(b, Day4{filePath: "inputs/day4/problem.txt"})
}
