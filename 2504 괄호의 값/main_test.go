package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 28, solve("(()[[]])([])"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 0, solve("([)]"))
	assert.Equal(t, 0, solve("(()()[]"))
}

func Test3(t *testing.T) {
	assert.Equal(t, 10, solve("([]())"))
}

func Test4(t *testing.T) {
	assert.Equal(t, 0, solve(""))
	assert.Equal(t, 0, solve(" "))
	assert.Equal(t, 0, solve("{}"))
}
