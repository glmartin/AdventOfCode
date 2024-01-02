package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var inputFileName = "2023/Day3/input"
var symbols = ""

type PossiblePartNumber struct {
	Value      string
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
	possibleParts := findPossibleNumbers(fileContents)
	numbers, err := findPartNumbers(fileContents, possibleParts)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, i := range numbers {
		total += i
	}
	return total, nil
}

func findPossibleNumbers(fileContents []string) []PossiblePartNumber {
	var idRegex = regexp.MustCompile("[0-9]+")
	numbers := make([]PossiblePartNumber, 0)
	for lineNumber, line := range fileContents {

		numStrs := idRegex.FindAllString(line, -1)

		for _, numStr := range numStrs {
			// find the index in the string of the number
			regexStr := numStr + "+"
			idxRegex := regexp.MustCompile(regexStr)
			indicies := idxRegex.FindAllStringIndex(line, -1)
			for _, idx := range indicies {
				possNum := PossiblePartNumber{
					LineNumber: lineNumber,
					Value:      numStr,
					Index:      idx[0],
				}
				numbers = append(numbers, possNum)
			}
		}
	}
	return numbers
}

func findPartNumbers(fileContents []string, possibleNumbers []PossiblePartNumber) ([]int, error) {
	numbers := make([]int, 0)

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

		if isPartNumbers(possNum, line, lineBefore, lineAfter) {
			partNum, err := strconv.Atoi(possNum.Value)
			if err != nil {
				return numbers, err
			}
			numbers = append(numbers, partNum)
		}
	}

	return numbers, nil
}

func isPartNumbers(possibleNumber PossiblePartNumber, line string, lineBefore string, lineAfter string) bool {

	lowestIndex := 0
	highestIndex := len(line)

	if possibleNumber.Index > 0 {
		lowestIndex = possibleNumber.Index - 1
	}

	if (possibleNumber.Index + len(possibleNumber.Value)) < len(line) {
		highestIndex = possibleNumber.Index + len(possibleNumber.Value) + 1
	}

	if possibleNumber.Index > 0 {
		// check the character to the left
		if isSymbol(line[(possibleNumber.Index - 1):possibleNumber.Index]) {
			return true
		}
	}

	if (possibleNumber.Index + len(possibleNumber.Value)) < len(line) {
		// check the character to the right
		if isSymbol(line[(possibleNumber.Index + len(possibleNumber.Value)) : (possibleNumber.Index+len(possibleNumber.Value))+1]) {
			return true
		}
	}

	if lineBefore != "" {
		// check the line above
		subLine := lineBefore[lowestIndex:highestIndex]
		if containsSymbol(subLine) {
			return true
		}
	}

	if lineAfter != "" {
		// check the line above
		subLine := lineAfter[lowestIndex:highestIndex]
		if containsSymbol(subLine) {
			return true
		}
	}

	return false
}

func isSymbol(char string) bool {
	return strings.Contains(symbols, char)
}

func containsSymbol(val string) bool {
	for _, char := range val {
		if strings.ContainsRune(symbols, char) {
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
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("--- Part 2")
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Value: %d", value)
}
