package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyllisMonteiro/go-api-template/pkg/models"
	v1 "github.com/wyllisMonteiro/go-api-template/pkg/api/v1"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	v1.InitRoutes(r)

	err := models.ConnectToDB()
	if err != nil {
		log.Println(err)
	}

	defer models.DB.Close()

	log.Fatal(http.ListenAndServe(":9000", r))
}
