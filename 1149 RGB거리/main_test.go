package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 96, solve([][3]int{
		{26, 40, 83},
		{49, 60, 57},
		{13, 89, 99},
	}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 101, solve([][3]int{
		{100, 1, 100},
		{999, 1, 999},
	}))
}
