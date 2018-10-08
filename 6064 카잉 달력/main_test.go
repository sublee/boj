package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, 33, solve(10, 12, 3, 9))
	assert.Equal(t, -1, solve(10, 12, 7, 2))
	assert.Equal(t, 83, solve(13, 11, 5, 6))
}

func TestMore(t *testing.T) {
	assert.Equal(t, 1, solve(13, 11, 1, 1))
	assert.Equal(t, 11, solve(10, 12, 1, 11))
	assert.Equal(t, 13, solve(10, 12, 3, 1))
	assert.Equal(t, 60, solve(10, 12, 10, 12))
	assert.Equal(t, -1, solve(10, 12, 9, 12))
}

func TestLarge(t *testing.T) {
	assert.Equal(t, 1599960000, solve(39999, 40000, 39999, 40000))
}
