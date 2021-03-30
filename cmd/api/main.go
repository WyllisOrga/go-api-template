package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/wyllisMonteiro/go-api-template/pkg/api/v1"
	"github.com/wyllisMonteiro/go-api-template/pkg/models"
	"github.com/wyllisMonteiro/go-api-template/pkg/swagger"
)

// @title Swagger API Project
// @version 1.0
// @description API.
// @termsOfService http://swagger.io/terms/

// @contact.email wyllismonteiro@gmail.com

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	gin.ForceConsoleColor()
	r := gin.Default()

	swagger.GetAPIDoc(r)
	v1.InitRoutes(r)

	err := models.ConnectToDB()
	if err != nil {
		log.Println(err)
	}

	defer models.DB.Close()

	log.Fatal(http.ListenAndServe(":8080", r))
}
