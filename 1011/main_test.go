package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 3, solve(0, 3))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 3, solve(1, 5))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, 4, solve(45, 50))
}

func TestMaximum(t *testing.T) {
	assert.Equal(t, 1, solve(0, 1<<31))
}
