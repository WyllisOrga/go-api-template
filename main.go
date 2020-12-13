package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/go-api-template/models"
	"github.com/wyllisMonteiro/go-api-template/router"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)

	err := models.ConnectToDB()
	if err != nil {
		log.Println(err)
	}

	defer models.DB.Close()

	log.Fatal(http.ListenAndServe(":9000", r))
}
