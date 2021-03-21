package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wyllisMonteiro/go-api-template/pkg/api/v1/sample"
)

//InitRoutes Load handlers
func InitRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/")
	{
		sample.Routes(api)
	}
}