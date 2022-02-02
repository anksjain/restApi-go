package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("using custom middleware")
		c.Next()
	}
}
