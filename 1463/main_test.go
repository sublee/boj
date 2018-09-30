package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 1, solve(2))
}

func Test2(t *testing.T) {
	assert.Equal(t, 3, solve(10))
}

func Test3(t *testing.T) {
	assert.Equal(t, 0, solve(1))
}

func Test4(t *testing.T) {
	assert.Equal(t, 19, solve(1e6))
}
