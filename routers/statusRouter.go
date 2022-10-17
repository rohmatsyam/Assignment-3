package routers

import (
	"assignment3/controllers/statuses"

	"github.com/gin-gonic/gin"
)

func StatusRouter(router gin.Engine) *gin.Engine {
	router.GET("/", statuses.GetStatus)
	return &router
}
