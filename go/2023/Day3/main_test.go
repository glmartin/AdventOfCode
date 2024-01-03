package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fileContents []string
var err error

var input = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

var possiblePartNumbers = []PossiblePartNumber{
	{
		Value:      "467",
		LineNumber: 0,
		Index:      0,
	},
	{
		Value:      "114",
		LineNumber: 0,
		Index:      5,
	},
	{
		Value:      "35",
		LineNumber: 2,
		Index:      2,
	},
	{
		Value:      "633",
		LineNumber: 2,
		Index:      6,
	},
	{
		Value:      "617",
		LineNumber: 4,
		Index:      0,
	},
	{
		Value:      "58",
		LineNumber: 5,
		Index:      7,
	},
	{
		Value:      "592",
		LineNumber: 6,
		Index:      2,
	},
	{
		Value:      "755",
		LineNumber: 7,
		Index:      6,
	},
	{
		Value:      "664",
		LineNumber: 9,
		Index:      1,
	},
	{
		Value:      "598",
		LineNumber: 9,
		Index:      5,
	},
}

func TestMain(m *testing.M) {
	fileContents, err = ScanFile("input")
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestScanFile(t *testing.T) {
	assert.Equal(t, 140, len(fileContents))
}

func TestCollectSymbols(t *testing.T) {
	collectSymbols(fileContents)

	assert.Equal(t, "*@-+#%=/$&", symbols)
}

func TestFindPossibleNumbers(t *testing.T) {
	result := findPossibleNumbers(input)
	assert.Equal(t, possiblePartNumbers, result)
}

func TestFindPartNumbers(t *testing.T) {

	collectSymbols(input)
	expectedResult := []int{467, 35, 633, 617, 592, 755, 664, 598}
	result, err := findPartNumbers(input, possiblePartNumbers)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestIsPartNumbers(t *testing.T) {

	// test against the actual file contents
	collectSymbols(fileContents)
	possibleNumber := PossiblePartNumber{
		Value:      "35",
		LineNumber: 2,
		Index:      2,
	}
	line := input[possibleNumber.LineNumber]
	lineBefore := input[possibleNumber.LineNumber-1]
	lineAfter := input[possibleNumber.LineNumber+1]

	result := isPartNumbers(possibleNumber, line, lineBefore, lineAfter)

	assert.Equal(t, true, result)

	possibleNumber = PossiblePartNumber{
		Value:      "3555",
		LineNumber: 2,
		Index:      5,
	}
	line = ".....3555...."
	lineBefore = "............."
	lineAfter = "............."

	for _, c := range symbols {
		// test other symbols
		line = fmt.Sprintf(".....3555%s...", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		line = fmt.Sprintf("....%s3555....", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		line = fmt.Sprintf(".....3555.%s..", string(c))
		assert.Equal(t, false, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		line = fmt.Sprintf("...%s.3555....", string(c))
		assert.Equal(t, false, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		line = ".....3555...."
		lineBefore = fmt.Sprintf("...%s.........", string(c))
		assert.Equal(t, false, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf("....%s........", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf(".....%s.......", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf("......%s......", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf(".......%s.....", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf("........%s....", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf(".........%s...", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = fmt.Sprintf("..........%s..", string(c))
		assert.Equal(t, false, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineBefore = "............"

		lineAfter = fmt.Sprintf("...%s.........", string(c))
		assert.Equal(t, false, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf("....%s........", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf(".....%s.......", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf("......%s......", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf(".......%s.....", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf("........%s....", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf(".........%s...", string(c))
		assert.Equal(t, true, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

		lineAfter = fmt.Sprintf("..........%s..", string(c))
		assert.Equal(t, false, isPartNumbers(possibleNumber, line, lineBefore, lineAfter))

	}
}

func TestIsSymbol(t *testing.T) {
	fileContents, err = ScanFile("input")
	assert.NoError(t, err)

	collectSymbols(fileContents)

	assert.Equal(t, false, isSymbol("."))
	assert.Equal(t, true, isSymbol("*"))
	assert.Equal(t, true, isSymbol("/"))
	assert.Equal(t, false, isSymbol("5"))
	assert.Equal(t, true, isSymbol("="))
	assert.Equal(t, true, isSymbol("#"))
	assert.Equal(t, true, isSymbol("&"))
}

func TestContainsSymbol(t *testing.T) {
	fileContents, err = ScanFile("input")
	assert.NoError(t, err)

	collectSymbols(fileContents)

	inputMap := map[string]bool{
		"467..114..": false,
		"...*......": true,
		"..35..633.": false,
		"......#...": true,
		"617*......": true,
		".....+.58.": true,
		"..592.....": false,
		"......755.": false,
		"...$.*....": true,
		".664.598..": false,
	}

	for line, expectedBool := range inputMap {
		assert.Equalf(t, expectedBool, containsSymbol(line), "%s is incorrect for %s", expectedBool, line)
	}
}
