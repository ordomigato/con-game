package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueGameCode(t *testing.T) {
	var usedCodes = []string{"abcd"}
	var result string = generateUniqueGameCode(usedCodes)
	fmt.Println(result)
	assert.Len(t, result, 4)
	for _, code := range usedCodes {
		assert.NotEqual(t, result, code)
	}
}

// func TestCreate(t *testing.T) {
// 	var mockGameConfig = entity.GameConfig{}
// 	var mockGame = entity.Game{
// 		GameConfig: mockGameConfig,
// 	}
// 	result := New().Create(mockGameConfig)
// 	assert.Equal(t, result, mockGame)
// }
