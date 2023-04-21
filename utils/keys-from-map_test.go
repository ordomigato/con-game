package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKeysFromMap(t *testing.T) {
	var testMap map[string]string = map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}
	result := GetKeysFromMap(testMap)
	assert.Equal(t, result, []string{"1", "2", "3"})
}
