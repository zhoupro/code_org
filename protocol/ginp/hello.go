package ginp

import (
	"code_org/handler"
	"code_org/handler/dto"
	"context"

	"github.com/gin-gonic/gin"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		req := dto.ReqHello{}
		req.Name = c.Query("name")
		res, err := handler.Hello(ctx, req)
		if err != nil {
			c.JSON(500, res)
		} else {
			c.JSON(200, res)
		}
	}
}
