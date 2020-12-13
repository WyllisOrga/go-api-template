package controllers

import (
	"net/http"

	"github.com/wyllisMonteiro/go-api-template/services"
)

// Sample stores sample data
type Sample struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetSample returns json with sample
func GetSample(w http.ResponseWriter, req *http.Request) {
	sample := Sample{
		ID:   1,
		Name: "Wyllis",
	}
	services.WriteJSON(w, http.StatusOK, sample)
}
