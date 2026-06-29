package main

import (
	routerGroup "GoReviews/gin/routerGroup"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//client
	routerGroup.ClientRouters(router)

	//platform
	routerGroup.PlatformRouters(router)

	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
