package routers

import (
	statuses "assignment3/controllers/Statuses"

	"github.com/gin-gonic/gin"
)

func StatusRouter(router gin.Engine) *gin.Engine {
	router.GET("/", statuses.GetStatus)
	return &router
}
