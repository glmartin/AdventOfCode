package main

import (
	"bufio"
	"log"
	"os"
)

var inputFileName = "2023/Day4/input"

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
	return 0, nil
}

func FindResultPart2(fileContents []string) (int, error) {
	return 0, nil
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
