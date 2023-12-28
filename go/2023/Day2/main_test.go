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
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green": Game{
			ID: 1,
			Sets: []GameSet{
				GameSet{
					Red:  4,
					Blue: 3,
				},
				GameSet{
					Red:   1,
					Blue:  6,
					Green: 2,
				},
				GameSet{
					Green: 2,
				},
			},
		},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue": Game{
			ID: 2,
			Sets: []GameSet{
				GameSet{
					Green: 2,
					Blue:  1,
				},
				GameSet{
					Red:   1,
					Blue:  4,
					Green: 3,
				},
				GameSet{
					Blue:  1,
					Green: 1,
				},
			},
		},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": Game{
			ID: 3,
			Sets: []GameSet{
				GameSet{
					Red:   20,
					Green: 8,
					Blue:  6,
				},
				GameSet{
					Red:   4,
					Blue:  5,
					Green: 13,
				},
				GameSet{
					Red:   1,
					Green: 5,
				},
			},
		},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": Game{
			ID: 4,
			Sets: []GameSet{
				GameSet{
					Red:   3,
					Green: 1,
					Blue:  6,
				},
				GameSet{
					Red:   6,
					Green: 3,
				},
				GameSet{
					Red:   14,
					Blue:  15,
					Green: 3,
				},
			},
		},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green": Game{
			ID: 5,
			Sets: []GameSet{
				GameSet{
					Red:   6,
					Green: 3,
					Blue:  1,
				},
				GameSet{
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
