package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wyllisMonteiro/go-api-template/pkg/api/v1/auth"
	"github.com/wyllisMonteiro/go-api-template/pkg/api/v1/sample"
	"github.com/wyllisMonteiro/go-api-template/pkg/jwt_auth"
)

//InitRoutes Load handlers
func InitRoutes(r *gin.Engine) {
	jwtAuth, err := jwt_auth.NewAuth()
	if err != nil {
		log.Println(err)
	}

	api := r.Group("/api/v1/")
	{
		auth.Routes(api, jwtAuth)
		sample.Routes(api, jwtAuth)
	}
}
