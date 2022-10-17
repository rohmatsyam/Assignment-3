package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")
	router.LoadHTMLFiles("templates/index.html")

	StatusRouter(*router)

	return router
}
