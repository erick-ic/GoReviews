package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageController struct{}

func (mc *MessageController) Success(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success!")
}
