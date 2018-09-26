package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, solve(0, 3))
	assert.Equal(t, 3, solve(1, 5))
	assert.Equal(t, 4, solve(45, 50))
}
