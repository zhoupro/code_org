package main

import (
	"code_org/protocol/ginp"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", ginp.Hello())
	router.Run(":8080")
}
