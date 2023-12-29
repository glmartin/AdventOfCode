package main

import (
	"bufio"
	"errors"
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

	// parse the game ID
	gameIdMatches := idRegex.FindSubmatch([]byte(lineparts[0]))
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

func isPossible(game Game, red int, green int, blue int) bool {
	for _, gameSet := range game.Sets {
		if red < gameSet.Red || blue < gameSet.Blue || green < gameSet.Green {
			return false
		}
	}

	return true
}

func checkGames(games []Game, red int, green int, blue int) int {
	total := 0

	for _, g := range games {
		if isPossible(g, red, green, blue) {
			total += g.ID
		}
	}
	return total

}

func findMinSet(game Game) GameSet {
	minSet := GameSet{}
	for _, gameSet := range game.Sets {
		if gameSet.Red > minSet.Red {
			minSet.Red = gameSet.Red
		}
		if gameSet.Blue > minSet.Blue {
			minSet.Blue = gameSet.Blue
		}
		if gameSet.Green > minSet.Green {
			minSet.Green = gameSet.Green
		}
	}

	return minSet
}

func findMinSetPower(game Game) int {
	gs := findMinSet(game)

	power := gs.Red
	power = power * gs.Blue
	power = power * gs.Green
	return power

}

func findMinSetsTotalPower(games []Game) int {
	total := 0

	for _, g := range games {
		power := findMinSetPower(g)
		total += power
	}
	return total

}

func parseGames(fileContents []string) ([]Game, error) {
	games := make([]Game, 0)
	for _, line := range fileContents {

		g, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 && len(args) != 4 {
		log.Fatal(errors.New(`Usage: 
		runme --part1 <red count> <green count> <blue count>
		runme --part2`))
	}

	fileContents, err := ScanFile(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGames(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	part := args[0]

	if part == "--part1" {
		var argInts = []int{}

		for _, i := range args[1:4] {
			j, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			argInts = append(argInts, j)
		}

		total := checkGames(games, argInts[0], argInts[1], argInts[2])
		log.Printf("Total: %v", total)
	} else {
		total := findMinSetsTotalPower(games)
		log.Printf("Total Power: %v", total)
	}
}
