package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Game 77: 3 green, 5 red, 8 blue; 14 red, 15 green; 14 green, 1 blue, 2 red
type Game struct {
	ID   int
	Sets []GameSet
}

type GameSet struct {
	Red   int
	Blue  int
	Green int
}

var inputFileName = "input"

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

func parseLine(line string) (Game, error) {

	g := Game{}

	var idRegex = regexp.MustCompile("[0-9]+")

	// Game 77: 3 green, 5 red, 8 blue; 14 red, 15 green; 14 green, 1 blue, 2 red
	lineparts := strings.Split(line, ":")

	fmt.Printf("%s  \n", lineparts[0])

	// parse the game ID
	gameIdMatches := idRegex.FindSubmatch([]byte(lineparts[0]))
	fmt.Printf("%s  \n", gameIdMatches)
	id, err := strconv.Atoi(string(gameIdMatches[0]))
	if err != nil {
		return g, err
	}

	g.ID = id

	// parse the sets
	gameSets := make([]GameSet, 0)
	sets := strings.Split(lineparts[1], ";")

	for _, s := range sets {

		gs := GameSet{}
		setSplit := strings.Split(s, ",")

		for _, setItem := range setSplit {

			var countRegex = regexp.MustCompile("[0-9]+")

			// parse the count
			fmt.Println(setItem)
			countMatches := countRegex.FindSubmatch([]byte(setItem))
			count, err := strconv.Atoi(string(countMatches[0]))
			if err != nil {
				return g, err
			}

			if strings.Contains(setItem, "red") {
				gs.Red = count
			} else if strings.Contains(setItem, "blue") {
				gs.Blue = count
			} else if strings.Contains(setItem, "green") {
				gs.Green = count
			}
		}

		gameSets = append(gameSets, gs)

	}

	g.Sets = gameSets

	return g, nil

}

func ParseGames(fileContents []string) ([]*Game, error) {

	for _, line := range fileContents {
		_, err := parseLine(line)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func main() {
	log.Println("BEGIN")
	fileContents, err := ScanFile(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	ParseGames(fileContents)
}
