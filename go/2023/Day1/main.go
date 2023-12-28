package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var inputFileName = "input"

var pt1Regex = regexp.MustCompile("\\d")

var pt2FirstRegex = regexp.MustCompile("([1-9]|zero|one|two|three|four|five|six|seven|eight|nine)")
var pt2LastRegex = regexp.MustCompile(".*([1-9]|zero|one|two|three|four|five|six|seven|eight|nine)")

func ParseInts(line string) (int, error) {
	matches := pt2FirstRegex.FindSubmatch([]byte(line))
	fstVal := toDigit(string(matches[1]))

	matches = pt2LastRegex.FindSubmatch([]byte(line))
	lstVal := toDigit(string(matches[1]))

	valStr := fstVal + lstVal
	i, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func toDigit(s string) string {
	switch s {
	case "zero":
		return "0"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
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

func FindResultPart1(fileContents []string) (int, error) {
	var value = 0
	// loop line by line to get the first and last digit of each line
	for _, line := range fileContents {
		numStrs := pt1Regex.FindAllString(line, -1)

		fstVal := numStrs[0]
		lstVal := numStrs[len(numStrs)-1]
		valStr := fstVal + lstVal
		i, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, err
		}

		// add the value
		value = value + i
	}
	return value, nil
}

func FindResultPart2(fileContents []string) (int, error) {
	var value = 0
	// loop line by line to get the first and last digit of each line
	for _, line := range fileContents {

		i, err := ParseInts(line)
		if err != nil {
			return 0, err
		}

		// add the value
		value += i
	}
	return value, nil
}

func main() {

	fileContents, err := ScanFile(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	value := 0

	args := os.Args[1:]
	if args[0] == "part1" {
		value, err = FindResultPart1(fileContents)
	} else {
		value, err = FindResultPart2(fileContents)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Value: %d", value)
}
