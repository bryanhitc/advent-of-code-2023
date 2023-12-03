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

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

func solveProblem(problem Problem) {
	start := time.Now()
	lines, err := readLines(problem.FilePath())
	handleErr(err)

	p1Start := time.Now()
	p1Val, err := problem.Part1(lines)
	p1Duration := time.Now().Sub(p1Start)
	handleErr(err)

	p2Start := time.Now()
	p2Val, err := problem.Part2(lines)
	p2Duration := time.Now().Sub(p2Start)
	handleErr(err)

	end := time.Now()
	fmt.Printf("Finished %s in %s: P1=%d (%s) P2=%d (%s)\n",
		problem.Name(),
		end.Sub(start),
		p1Val,
		p1Duration,
		p2Val,
		p2Duration,
	)
}

func main() {
	solveProblem(Day01{filePath: "inputs/day01/problem.txt"})
	solveProblem(Day02{filePath: "inputs/day02/problem.txt"})
}
