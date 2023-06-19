package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	engine.Run(":8080")
}
