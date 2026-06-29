package routerGroup

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientRouters(router *gin.Engine) {
	clientRouters := router.Group("/client")
	{
		clientRouters.GET("/index", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "this is client index")
		})
		clientRouters.GET("/list", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "this is client list")
		})
	}
}
