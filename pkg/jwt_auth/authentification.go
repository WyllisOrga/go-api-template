package jwt_auth

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
	"github.com/wyllisMonteiro/go-api-template/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User
type Role = models.Role

//NewAuth Create new auth context
func NewAuth() (jwt.Auth, error) {
	return jwt.New(jwt.Auth{
		SecretKey: []byte("S3CR3TK3Y733T"),
		Authenticator: func(c *gin.Context) (jwt.MapClaims, error) {
			var req models.RequestLogin

			if err := c.ShouldBind(&req); err != nil {
				return nil, jwt.ErrorAuthenticationFailed
			}

			u, err := models.GetUser(req.Username)

			if err != nil {
				return nil, jwt.ErrorUserNotFound
			}

			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
				return nil, jwt.ErrorAuthenticationFailed
			}

			return jwt.MapClaims{
				"username": u.Username,
				"role":     u.Role,
			}, nil
		},
		UserFetcher: func(c *gin.Context, claims jwt.MapClaims) (interface{}, error) {
			username, ok := claims["username"].(string)
			if !ok {
				return nil, nil
			}
			u, err := models.GetUser(username)
			if err != nil {
				return nil, nil
			}
			return u, nil
		},
	})
}

//Operator is user is admin
func Operator(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsOperator()
	})
}

//Admin is user is admin
func Admin(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsAdmin()
	})
}

//SystemAdmin is user is system admin
func SystemAdmin(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsSystemAdmin()
	})
}

//role get the role of the user
func role(claims jwt.MapClaims) Role {
	return Role(claims["role"].(float64))
}
