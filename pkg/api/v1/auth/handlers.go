package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyllisMonteiro/go-api-template/pkg/models"
)

type requestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//Register routes for creating account
func Register(c *gin.Context) {
	var req models.RequestRegister

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err := models.CreateUser(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account_created",
	})
}
