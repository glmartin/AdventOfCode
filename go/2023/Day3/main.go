package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

var inputFileName = "2023/Day3/input"
var symbols = ""

type PossiblePartNumber struct {
	Value      string
	LineNumber int
	Index      int
	AdjSymbol  AdjacentSymbol
}

type AdjacentSymbol struct {
	Symbol     byte
	LineNumber int
	Index      int
}

func ScanFile(inputFile string) ([]string, error) {
	// open the inoput file
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the file into the slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func collectSymbols(fileContents []string) {
	for _, line := range fileContents {
		for _, char := range line {
			if !unicode.IsDigit(char) && char != '.' && !strings.ContainsRune(symbols, char) {
				symbols += string(char)
			}
		}
	}
}

func FindResultPart1(fileContents []string) (int, error) {
	partNumbers, err := findParts(fileContents)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, part := range partNumbers {
		partInt, err := strconv.Atoi(part.Value)
		if err != nil {
			return 0, err
		}

		total += partInt
	}
	return total, nil
}

func FindResultPart2(fileContents []string) (int, error) {
	// In this case, we'll rerun a similar search to part 1, but only for number adjacent to the astrerisk symbol
	symbols = "*"

	partNumbers, err := findParts(fileContents)
	if err != nil {
		return 0, err
	}

	total := 0

	// loop through each partnumber in a nested loop, finding where they have a common adjacent asterisk
	alreadyChecked := make([]PossiblePartNumber, 0)
	for _, part := range partNumbers {
		alreadyChecked = append(alreadyChecked, part)
		for _, otherPart := range partNumbers {
			if part != otherPart && !slices.Contains(alreadyChecked, otherPart) {
				// do the 2 parts share the same adjacent symbol?
				if part.AdjSymbol == otherPart.AdjSymbol {
					partInt, err := strconv.Atoi(part.Value)
					if err != nil {
						return 0, err
					}
					otherPartInt, err := strconv.Atoi(otherPart.Value)
					if err != nil {
						return 0, err
					}			
					total = total + (partInt * otherPartInt)
				}
			}
		}
	}
	return total, nil
}

func findParts(fileContents []string) ([]PossiblePartNumber, error) {
	possibleParts := findPossibleNumbers(fileContents)
	return findPartNumbers(fileContents, possibleParts)
}

func findPossibleNumbers(fileContents []string) []PossiblePartNumber {
	var idRegex = regexp.MustCompile("[0-9]+")
	numbers := make([]PossiblePartNumber, 0)
	for lineNumber, line := range fileContents {

		numStrs := idRegex.FindAllString(line, -1)
		indicies := idRegex.FindAllStringIndex(line, -1)

		for i, idx := range indicies {
			possNum := PossiblePartNumber{
				LineNumber: lineNumber,
				Value:      numStrs[i],
				Index:      idx[0],
			}
			numbers = append(numbers, possNum)
		}
	}
	return numbers
}

func findPartNumbers(fileContents []string, possibleNumbers []PossiblePartNumber) ([]PossiblePartNumber, error) {
	numbers := make([]PossiblePartNumber, 0)

	// look through all of the possible numbers, and only same the numbers that are adjacent to a symbol (other than period).
	// the code will need to look to the left, right of the string, as well as the line above and below
	for _, possNum := range possibleNumbers {
		line := fileContents[possNum.LineNumber]
		lineBefore := ""
		lineAfter := ""

		if possNum.LineNumber > 0 {
			lineBefore = fileContents[possNum.LineNumber-1]
		}
		if possNum.LineNumber < len(fileContents)-1 {
			lineAfter = fileContents[possNum.LineNumber+1]
		}

		if isPart, symbol, symbolLine, symbolIdx := isPartNumber(possNum, line, lineBefore, lineAfter); isPart {
			// save the adjacent symbol
			adjSymbol := AdjacentSymbol{
				Symbol:     symbol,
				LineNumber: symbolLine,
				Index:      symbolIdx,
			}
			possNum.AdjSymbol = adjSymbol
			numbers = append(numbers, possNum)
		}
	}

	return numbers, nil
}

// This function will check the strings before, after, above, and below a possible number for symbol.
// it returns a true if it is a part number, along with the symbol and symbol location of the adjacent symbol
func isPartNumber(possibleNumber PossiblePartNumber, line string, lineBefore string, lineAfter string) (bool, byte, int, int) {

	lowestIndex := 0
	highestIndex := len(line)
	endOfNumberIdx := possibleNumber.Index + len(possibleNumber.Value)

	if possibleNumber.Index > 0 {
		lowestIndex = possibleNumber.Index - 1
	}

	if endOfNumberIdx < len(line) {
		highestIndex = endOfNumberIdx + 1
	}

	if possibleNumber.Index > 0 {
		// check the character to the left
		if isSymbol(line[(possibleNumber.Index - 1):possibleNumber.Index]) {
			symbol := line[(possibleNumber.Index - 1)]
			return true, symbol, possibleNumber.LineNumber, possibleNumber.Index - 1
		}
	}

	if endOfNumberIdx < len(line) {
		// check the character to the right
		if isSymbol(line[endOfNumberIdx : endOfNumberIdx + 1]) {
			symbol := line[endOfNumberIdx]
			return true, symbol, possibleNumber.LineNumber, endOfNumberIdx
		}
	}

	if lineBefore != "" {
		// check the line above
		subLine := lineBefore[lowestIndex:highestIndex]
		if containsSymbol(subLine) {
			symbIdx := strings.Index(subLine, symbols)
			var symbol byte = 0
			if symbIdx > -1 {
				symbol = subLine[symbIdx]
			}
			return true, symbol, possibleNumber.LineNumber - 1, symbIdx + lowestIndex
		}
	}

	if lineAfter != "" {
		// check the line below
		subLine := lineAfter[lowestIndex:highestIndex]
		if containsSymbol(subLine) {
			symbIdx := strings.Index(subLine, symbols)
			var symbol byte = 0
			if symbIdx > -1 {
				symbol = subLine[symbIdx]
			}
			return true, symbol, possibleNumber.LineNumber + 1, symbIdx + lowestIndex
		}
	}

	return false, 0, 0, 0
}

func isSymbol(char string) bool {
	return strings.Contains(symbols, char)
}

func containsSymbol(val string) bool {
	for _, char := range val {
		if isSymbol(string(char)) {
			return true
		}
	}
	return false
}

func main() {
	log.Println("**** 2023 - Day 3 ****")
	fileContents, err := ScanFile(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	collectSymbols(fileContents)
	value := 0

	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal(`Usage: 
		runme --part1
		runme --part2`)
	}
	if args[0] == "--part1" {
		log.Println("--- Part 1")
		value, err = FindResultPart1(fileContents)
	} else {
		log.Println("--- Part 2")
		value, err = FindResultPart2(fileContents)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Value: %d", value)
}
