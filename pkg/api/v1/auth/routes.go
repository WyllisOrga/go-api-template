package auth

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

//Routes for authentification
func Routes(r *gin.RouterGroup, jwtAuth jwt.Auth) {
	r.Use(jwt.ErrorHandler)
	r.POST("/login", jwtAuth.Authenticate)
	r.POST("/register", Register)
}
