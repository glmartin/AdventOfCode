package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindResultPart1(t *testing.T) {
	scanner, err := ScanFile(inputFileName)
	assert.NoError(t, err)

	value, err := FindResultPart1(scanner)

	assert.NoError(t, err)

	var expected = 53194
	assert.Equal(t, expected, value)
}

func TestFindResultPart2(t *testing.T) {
	scanner, err := ScanFile(inputFileName)
	assert.NoError(t, err)

	value, err := FindResultPart2(scanner)

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
