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
