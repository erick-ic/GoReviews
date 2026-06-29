package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndex(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is admin index")

}

func UserList(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is admin list")
}

//type UserController struct {
//	MessageController //继承父级结构体
//}
//
//func (uc *UserController) Index(ctx *gin.Context) {
//	ctx.String(http.StatusOK, "this is admin index")
//
//}
//
//func (uc *UserController) List(ctx *gin.Context) {
//	ctx.String(http.StatusOK, "this is admin list")
//}

type UserController struct {
	MessageController //继承父级结构体
}

func (uc *UserController) Index(ctx *gin.Context) {
	uc.Success(ctx)

}

func (uc *UserController) List(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is admin list")
	uc.Success(ctx)
}
