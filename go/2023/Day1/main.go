package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func GetCodeFromFile(inputFile string) (int, error) {
	var value = 0

	// open the inoput file
	readFile, err := os.Open("input")
	if err != nil {
		return 0, err
	}
	defer readFile.Close()

	// read the file
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// loop line by line to get the first and last digit of each line
	for fileScanner.Scan() {
		re := regexp.MustCompile("\\d")
		numStrs := re.FindAllString(fileScanner.Text(), -1)

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

func main() {
	value, err := GetCodeFromFile("input")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Value: %d", value)
}
