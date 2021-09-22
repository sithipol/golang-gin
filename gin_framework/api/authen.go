package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginEndpoint(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
