package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize the router
	router := gin.Default()

	//Initialize Routes
	InitializeRoutes(router)

	//Run the Server
	router.Run(":4100")
}
