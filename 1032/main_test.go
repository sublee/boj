package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, "config????", solve([]string{"config.sys", "config.inf", "configures"}))
}
