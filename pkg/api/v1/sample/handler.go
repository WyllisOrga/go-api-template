package sample

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//Sample stores sample data
type Sample struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//GetSample returns json with sample
func GetSample(c *gin.Context) {
	sample := Sample{
		ID:   1,
		Name: "Wyllis",
	}

	c.JSON(http.StatusOK, sample)
}
