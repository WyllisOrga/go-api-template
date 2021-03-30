package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sample stores sample data
type Sample struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetSample Returns json with sample
// @Summary Show Sample
// @Description Get all samples
// @Tags sample
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} Sample
// @Failure 500 {object} httputil.HTTPError
// @Router / [get]
func GetSample(c *gin.Context) {
	sample := Sample{
		ID:   1,
		Name: "Wyllis",
	}

	c.JSON(http.StatusOK, sample)
}
