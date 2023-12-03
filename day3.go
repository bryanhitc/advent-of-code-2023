package main

type Day3 struct {
	filePath string
}

func (Day3) Name() string {
	return "day3"
}
func (day Day3) FilePath() string {
	return day.filePath
}

// lol... there's no built-in abs function for integers - only floats.
// I assume lack of generics was the original reason why?? jeez.

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isDigit(char byte) bool {
	return char >= byte('0') && char <= byte('9')
}

type Empty struct{}

type PartNumber struct {
	startIndex int
	endIndex   int
	num        int
	visited    bool
}

func (p PartNumber) IsAdjacent(i int) bool {
	return Abs(i-p.startIndex) <= 1 || Abs(i-p.endIndex) <= 1
}

type Symbol struct {
	lineIndex int
	colIndex  int
	symbol    byte
}

func parse(lines []string) ([][]PartNumber, []Symbol) {
	// Represents all numbers given a line index. Sorted by `startIndex` and `endIndex`.
	lineToPartNums := make([][]PartNumber, len(lines))
	symbols := make([]Symbol, 0)

	// Unicode makes the parsing code more complicated (can't manually increment iterator...), so
	// I'm just going to iterate over the bytes instead.
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := line[j]

			// garbage
			if char == byte('.') {
				continue
			}

			// symbol
			if !isDigit(char) {
				symbols = append(symbols, Symbol{
					lineIndex: i,
					colIndex:  j,
					symbol:    char,
				})
				continue
			}

			// digit
			startIndex := j
			num := 0
			for isDigit(char) {
				num *= 10
				num += int(char - byte('0'))
				j++
				if j >= len(line) {
					break
				}
				char = line[j]
			}

			// current character isn't a digit. let's rewind...
			j--
			endIndex := j
			part := PartNumber{
				startIndex: startIndex,
				endIndex:   endIndex,
				num:        num,
				visited:    false,
			}
			lineToPartNums[i] = append(lineToPartNums[i], part)
		}
	}
	return lineToPartNums, symbols
}

// two-pass algorithm:
//  1. parse all numbers and symbols into a sane, searchable format.
//  2. visit all symbols and determine which numbers "belong" to them.
func (Day3) Part1(lines []string) (int, error) {
	lineToPartNums, symbols := parse(lines)
	total := 0
	for _, symbol := range symbols {
		checkLine := func(s Symbol, line []PartNumber) {
			// yes, binary search here for finding part numbers (aka "intervals") given an index is best
			// here, but given that the input lines are relatively short, it's probably slower than this
			// cache-efficient brute-force approach.
			//
			// more generally, using an interval tree can work.
			for i := 0; i < len(line); i++ {
				partNum := &line[i]
				if partNum.visited || !partNum.IsAdjacent(symbol.colIndex) {
					continue
				}
				total += partNum.num
				partNum.visited = true
			}
		}

		// check symbol's line
		checkLine(symbol, lineToPartNums[symbol.lineIndex])

		// check above
		if symbol.lineIndex > 0 {
			checkLine(symbol, lineToPartNums[symbol.lineIndex-1])
		}

		// check below
		if symbol.lineIndex < len(lines)-1 {
			checkLine(symbol, lineToPartNums[symbol.lineIndex+1])
		}
	}

	return total, nil
}

// two-pass algorithm:
//  1. parse all numbers and symbols into a sane, searchable format.
//  2. visit all "gears" and determine whether they have exactly 2 numbers.
func (Day3) Part2(lines []string) (int, error) {
	lineToPartNums, symbols := parse(lines)
	gearNumbers := make([]*PartNumber, 2)
	total := 0
	for _, symbol := range symbols {
		if symbol.symbol != byte('*') {
			continue
		}

		numParts := 0
		exceedsLimit := false
		checkLine := func(s Symbol, line []PartNumber) {
			// yes, binary search here for finding part numbers (aka "intervals") given an index is best
			// here, but given that the input lines are relatively short, it's probably slower than this
			// cache-efficient brute-force approach.
			//
			// more generally, using an interval tree can work.
			if exceedsLimit {
				return
			}
			for i := 0; i < len(line); i++ {
				partNum := &line[i]
				if partNum.visited || !partNum.IsAdjacent(symbol.colIndex) {
					continue
				}
				partNum.visited = true
				if numParts >= 2 {
					exceedsLimit = true
					return
				}
				gearNumbers[numParts] = partNum
				numParts++
			}
		}

		// check symbol's line
		checkLine(symbol, lineToPartNums[symbol.lineIndex])

		// check above
		if symbol.lineIndex > 0 {
			checkLine(symbol, lineToPartNums[symbol.lineIndex-1])
		}

		// check below
		if symbol.lineIndex < len(lines)-1 {
			checkLine(symbol, lineToPartNums[symbol.lineIndex+1])
		}

		if !exceedsLimit && numParts == 2 {
			total += gearNumbers[0].num * gearNumbers[1].num
		}
		clear(gearNumbers)
	}
	return total, nil
}
