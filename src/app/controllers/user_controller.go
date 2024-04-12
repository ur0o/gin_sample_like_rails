package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserShow(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"user_id": user_id,
		"name": "fuga",
	})
}
