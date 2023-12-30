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

	assert.Equal(t, 140, len(fileContents))
}