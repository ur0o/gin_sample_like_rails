package config

import (
	"github.com/gin-gonic/gin"

	"gin_sample/app/controllers"
)

func SetRouting(e *gin.Engine) {
	users := e.Group("/users")
	{
		users.GET("/:id", controllers.UserShow)
	}
}