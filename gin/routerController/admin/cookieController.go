package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CookieController struct{}

func (cc *CookieController) CreateCookie(ctx *gin.Context) {
	//设置cookie 1h过期
	ctx.SetCookie("username", "Jack", 3600, "/", "localhost", false, true)
	ctx.String(http.StatusOK, "create cookie success")
}

func (cc *CookieController) DeleteCookie(ctx *gin.Context) {
	//删除cookie，即过期时间改为负数
	ctx.SetCookie("username", "Jack", -1, "/", "localhost", false, true)
	ctx.String(http.StatusOK, "delete cookie success")

}

func (cc *CookieController) EditCookie(ctx *gin.Context) {
	ctx.SetCookie("username", "Mary", 3600, "/", "localhost", false, true)
	ctx.String(http.StatusOK, "edit cookie success")

}

func (cc *CookieController) GetCookie(ctx *gin.Context) {
	//获取cookie
	username, _ := ctx.Cookie("username")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
	})
}
