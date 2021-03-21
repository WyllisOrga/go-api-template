package sample

import (
	"github.com/gin-gonic/gin"
)

//Routes All routes for sample
func Routes(r *gin.RouterGroup) {
	r.GET("", GetSample)
}