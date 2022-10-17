package statuses

import (
	"assignment3/models"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	r1 := rand.Intn(100)
	r2 := rand.Intn(100)
	status := &models.Status{
		Water: uint(r1),
		Wind:  uint(r2),
	}
	data, _ := json.Marshal(status)
	var hasil models.Status
	json.Unmarshal(data, &hasil)

	var statusStr string
	if hasil.Water < 5 || hasil.Wind < 6 {
		statusStr = "Aman"
	} else if hasil.Water < 8 || hasil.Wind < 15 {
		statusStr = "Siaga"
	} else {
		statusStr = "Bahaya"
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"wind":   hasil.Wind,
		"water":  hasil.Water,
		"status": statusStr,
	})
}
