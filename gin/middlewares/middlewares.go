package middlewares

import "github.com/gin-gonic/gin"

func SetMiddleware(ctx *gin.Context) {
	ctx.Set("username", "Jack")
}
