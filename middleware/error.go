package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			fmt.Printf("Error: %s\n", err)
		}

		c.JSON(-1, c.Errors[0])
	}
}
