package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		fmt.Printf("Error: %s\n", err)
	}
}
