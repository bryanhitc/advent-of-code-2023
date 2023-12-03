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

type SymbolLocation struct {
	lineIndex int
	colIndex  int
	symbol    byte
}

func parse(lines []string) ([][]PartNumber, []SymbolLocation) {
	// Represents all numbers given a line index. Sorted by `startIndex` and `endIndex`.
	lineToPartNums := make([][]PartNumber, len(lines))
	symbolLocations := make([]SymbolLocation, 0)

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
				symbolLocations = append(symbolLocations, SymbolLocation{
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
	return lineToPartNums, symbolLocations
}

// two-pass algorithm:
//  1. parse all numbers into a sane, searchable format. also make note of symbol indicies.
//  2. visit all symbols and determine which numbers "belong" to them.
func (Day3) Part1(lines []string) (int, error) {
	lineToPartNums, symbolLocations := parse(lines)
	total := 0

	// iterate through symbols to find nearby numbers!
	for _, symbolLocation := range symbolLocations {
		checkLine := func(s SymbolLocation, line []PartNumber) {
			// yes, binary search here for finding part numbers (aka "intervals") given an index is best
			// here, but given that the input lines are relatively short, it's doesn't matter + is prob slower.
			//
			// more generally, using an interval tree can work.
			for i := 0; i < len(line); i++ {
				partNum := &line[i]
				if partNum.visited || !partNum.IsAdjacent(symbolLocation.colIndex) {
					continue
				}
				total += partNum.num
				partNum.visited = true
			}
		}

		// check symbol's line
		checkLine(symbolLocation, lineToPartNums[symbolLocation.lineIndex])

		// check above
		if symbolLocation.lineIndex > 0 {
			checkLine(symbolLocation, lineToPartNums[symbolLocation.lineIndex-1])
		}

		// check below
		if symbolLocation.lineIndex < len(lines)-1 {
			checkLine(symbolLocation, lineToPartNums[symbolLocation.lineIndex+1])
		}
	}

	return total, nil
}
func (Day3) Part2(lines []string) (int, error) {
	lineToPartNums, symbolLocations := parse(lines)
	gearNumbers := make([]*PartNumber, 2)
	total := 0
	// iterate through symbols to find nearby numbers!
	for _, symbolLocation := range symbolLocations {
		if symbolLocation.symbol != byte('*') {
			continue
		}

		numParts := 0
		exceedsLimit := false
		checkLine := func(s SymbolLocation, line []PartNumber) {
			// yes, binary search here for finding part numbers (aka "intervals") given an index is best
			// here, but given that the input lines are relatively short, it's doesn't matter + is prob slower.
			//
			// more generally, using an interval tree can work.
			if exceedsLimit {
				return
			}
			for i := 0; i < len(line); i++ {
				partNum := &line[i]
				if partNum.visited || !partNum.IsAdjacent(symbolLocation.colIndex) {
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
		checkLine(symbolLocation, lineToPartNums[symbolLocation.lineIndex])

		// check above
		if symbolLocation.lineIndex > 0 {
			checkLine(symbolLocation, lineToPartNums[symbolLocation.lineIndex-1])
		}

		// check below
		if symbolLocation.lineIndex < len(lines)-1 {
			checkLine(symbolLocation, lineToPartNums[symbolLocation.lineIndex+1])
		}

		if !exceedsLimit && numParts == 2 {
			total += gearNumbers[0].num * gearNumbers[1].num
		}
		clear(gearNumbers)
	}
	return total, nil
}
