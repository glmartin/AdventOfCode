package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fileContents []string
var err error

func TestMain(m *testing.M) {
	fileContents, err = ScanFile("input")
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestScanFile(t *testing.T) {
	fileContents, err = ScanFile("input")
	assert.NoError(t, err)

	assert.Equal(t, 100, len(fileContents))
}

func TestParseLine(t *testing.T) {

	inputMap := map[string]Game{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green": {
			ID: 1,
			Sets: []GameSet{
				{
					Red:  4,
					Blue: 3,
				},
				{
					Red:   1,
					Blue:  6,
					Green: 2,
				},
				{
					Green: 2,
				},
			},
		},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue": {
			ID: 2,
			Sets: []GameSet{
				{
					Green: 2,
					Blue:  1,
				},
				{
					Red:   1,
					Blue:  4,
					Green: 3,
				},
				{
					Blue:  1,
					Green: 1,
				},
			},
		},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": {
			ID: 3,
			Sets: []GameSet{
				{
					Red:   20,
					Green: 8,
					Blue:  6,
				},
				{
					Red:   4,
					Blue:  5,
					Green: 13,
				},
				{
					Red:   1,
					Green: 5,
				},
			},
		},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": {
			ID: 4,
			Sets: []GameSet{
				{
					Red:   3,
					Green: 1,
					Blue:  6,
				},
				{
					Red:   6,
					Green: 3,
				},
				{
					Red:   14,
					Blue:  15,
					Green: 3,
				},
			},
		},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green": {
			ID: 5,
			Sets: []GameSet{
				{
					Red:   6,
					Green: 3,
					Blue:  1,
				},
				{
					Red:   1,
					Blue:  2,
					Green: 2,
				},
			},
		},
	}

	for line, expectedStruct := range inputMap {

		g, err := parseLine(line)
		assert.NoError(t, err)

		assert.Equal(t, expectedStruct, g)
	}
}

func TestIsPossible(t *testing.T) {
	// check using 12 red, 13 green, and 14 blue

	inputMap := map[string]bool{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   true,
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         true,
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": false,
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": false,
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   true,
	}

	for line, expectedBool := range inputMap {

		g, err := parseLine(line)
		assert.NoError(t, err)

		ip := isPossible(g, 12, 13, 14)

		assert.Equal(t, expectedBool, ip)
	}
}

func TestCheckGames(t *testing.T) {
	// check using 12 red, 13 green, and 14 blue

	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	games := make([]Game, 0)
	for _, line := range input {

		g, err := parseLine(line)
		assert.NoError(t, err)
		games = append(games, g)
	}

	assert.Equal(t, 8, checkGames(games, 12, 13, 14))
}

func TestFindMinSet(t *testing.T) {

	inputMap := map[string]GameSet{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green": {
			Red:   4,
			Blue:  6,
			Green: 2,
		},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue": {
			Red:   1,
			Blue:  4,
			Green: 3,
		},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": {
			Red:   20,
			Blue:  6,
			Green: 13,
		},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": {
			Red:   14,
			Blue:  15,
			Green: 3,
		},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green": {
			Red:   6,
			Blue:  2,
			Green: 3,
		},
	}

	for line, expectedStruct := range inputMap {

		g, err := parseLine(line)
		assert.NoError(t, err)

		gs := findMinSet(g)

		assert.Equal(t, expectedStruct, gs)
	}
}

func TestFindMinSetPower(t *testing.T) {

	inputMap := map[string]int{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   48,
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         12,
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": 1560,
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": 630,
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   36,
	}

	for line, expectedPower := range inputMap {

		g, err := parseLine(line)
		assert.NoError(t, err)

		power := findMinSetPower(g)

		assert.Equal(t, expectedPower, power)
	}
}

func TestFindMinSetsTotalPower(t *testing.T) {
	// check using 12 red, 13 green, and 14 blue

	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	games := make([]Game, 0)
	for _, line := range input {

		g, err := parseLine(line)
		assert.NoError(t, err)
		games = append(games, g)
	}

	assert.Equal(t, 2286, findMinSetsTotalPower(games))
}
