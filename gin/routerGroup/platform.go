package routerGroup

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlatformRouters(router *gin.Engine) {
	platformRouters := router.Group("/platform")
	{
		platformRouters.GET("/index", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "this is platform index")
		})
		platformRouters.GET("/list", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "this is platform list")
		})
	}
}
