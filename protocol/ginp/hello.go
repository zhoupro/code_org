package ginp

import "github.com/gin-gonic/gin"

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "hello world yes")
	}
}
