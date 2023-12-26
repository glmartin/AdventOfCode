package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	value, err := GetCodeFromFile("input")

	assert.NoError(t, err)

	var expected = 53194
	assert.Equal(t, expected, value)
}
