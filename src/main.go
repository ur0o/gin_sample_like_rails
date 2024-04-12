package main

import (
	"github.com/gin-gonic/gin"

	"gin_sample/config"
)

func main() {
	engine := gin.Default()

	config.SetRouting(engine)
	engine.Run(":8080")
}
