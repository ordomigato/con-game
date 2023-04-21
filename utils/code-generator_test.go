package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	result := GenerateCode(4)
	assert.Len(t, result, 4)
}
