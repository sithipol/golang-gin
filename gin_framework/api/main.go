package api

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login/:name", LoginEndpoint)
		// authenAPI.GET("/register",registerEndpoint)
	}

	uploadAPI := router.Group("/upload")
	{
		uploadAPI.GET("/", UploadFile)
		// authenAPI.GET("/register",registerEndpoint)
	}
}
