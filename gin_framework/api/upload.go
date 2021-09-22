package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": "ok",
	})
}
