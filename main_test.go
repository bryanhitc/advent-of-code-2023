package main

import (
	"testing"
)

func ProblemTestRunner(t *testing.T, problem Problem, expectedPart1 int, expectedPart2 int) {
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

func ProblemBenchmarkRunner(b *testing.B, problem Problem) {
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

func TestDay01(t *testing.T) {
	ProblemTestRunner(t, Day01{filePath: "inputs/day01/test01.txt"}, 142, 142)
	ProblemTestRunner(t, Day01{filePath: "inputs/day01/test02.txt"}, 209, 281)
	ProblemTestRunner(t, Day01{filePath: "inputs/day01/problem.txt"}, 54388, 53515)
}

func BenchmarkDay01(b *testing.B) {
	ProblemBenchmarkRunner(b, Day01{filePath: "inputs/day01/problem.txt"})
}
