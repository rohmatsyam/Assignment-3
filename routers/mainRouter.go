package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLFiles("templates/index.html")

	StatusRouter(*router)

	return router
}
