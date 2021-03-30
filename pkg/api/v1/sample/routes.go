package sample

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
	"github.com/wyllisMonteiro/go-api-template/pkg/jwt_auth"
)

//Routes All routes for sample
func Routes(r *gin.RouterGroup, jwtAuth jwt.Auth) {
	r.Use(jwt.ErrorHandler)
	r.GET("/", jwt_auth.Operator(jwtAuth), GetSample)
}
