package middleware

import "github.com/gin-gonic/gin"

func SetHeaders(c *gin.Context) {
	c.Request.Header.Add("Access-Control-Allow-Origin", "*")
	c.Next()
}
