package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Initialize the router
	router := gin.Default()

	//Initialize Routes
	InitializeRoutes(router)

	//Run the Server
	const port = ":4100"
	fmt.Println("🚀 Server running on http://localhost:4100")
	router.Run(port)

}
