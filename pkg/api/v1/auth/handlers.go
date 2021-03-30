package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyllisMonteiro/go-api-template/pkg/models"
)

type responseRegister struct {
	Message string `json:"message"`
}

// Register New account
// @Summary Create new account
// @Description Using JWT auth
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body models.RequestRegister true "Add account"
// @Success 200 {string} string "message"
// @Failure 500 {object} httputil.HTTPError
// @Router /register [post]
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

	c.JSON(http.StatusServiceUnavailable, gin.H{
		"message": "created",
	})
}

type requestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login An account
// @Summary Connect user to app
// @Description Using JWT auth (look headers for token)
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body requestLogin true "Log account"
// @Success 200 {string} string ""
// @Failure 500 {object} httputil.HTTPError
// @Router /login [post]
func Login(c *gin.Context) {
	// call jwt.Auth.Authenticate()
	// useless func, but help to make swagger doc
}