package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem interface {
	Name() string
	FilePath() string
	Part1(lines []string) (int, error)
	Part2(lines []string) (int, error)
}

func readLines(filePath string) ([]string, error) {
	var lines []string
	file, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func solve(problem Problem) {
	start := time.Now()
	lines, err := readLines(problem.FilePath())
	if err != nil {
		log.Fatal(err)
	}

	p1Start := time.Now()
	p1Val, err := problem.Part1(lines)
	p1Duration := time.Since(p1Start)
	if err != nil {
		log.Fatal(err)
	}

	p2Start := time.Now()
	p2Val, err := problem.Part2(lines)
	p2Duration := time.Since(p2Start)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Finished %s in %s: P1=%d (%s) P2=%d (%s)\n",
		problem.Name(),
		time.Since(start),
		p1Val,
		p1Duration,
		p2Val,
		p2Duration,
	)
}

func main() {
	solve(Day1{filePath: "inputs/day1/problem.txt"})
	solve(Day2{filePath: "inputs/day2/problem.txt"})
	solve(Day3{filePath: "inputs/day3/problem.txt"})
}
