package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSearchSlice(t *testing.T) {
	var result bool
	result = StringSearchSlice([]string{"hello"}, "world")
	assert.False(t, result)
	result = StringSearchSlice([]string{"hello"}, "hello")
	assert.True(t, result)
}
