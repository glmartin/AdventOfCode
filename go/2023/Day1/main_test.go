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

	assert.Equal(t, len(fileContents), 1000)
}

func TestFindResultPart1(t *testing.T) {
	value, err := FindResultPart1(fileContents)

	assert.NoError(t, err)

	var expected = 53194
	assert.Equal(t, expected, value)
}

func TestFindResultPart2(t *testing.T) {
	value, err := FindResultPart2(fileContents)

	assert.NoError(t, err)

	var expected = 54249
	assert.Equal(t, expected, value)
}

func TestParseInts(t *testing.T) {

	inputMap := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	total := 0
	for input, expectedInt := range inputMap {

		i, err := ParseInts(input)
		assert.NoError(t, err)

		assert.Equal(t, expectedInt, i)
		total += i
	}
	assert.Equal(t, 281, total)
}
