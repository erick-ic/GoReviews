package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct{}

func (nc *ProfileController) GetProfile(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	// 类型断言，Get()获取的是interface{}类型数据
	val, ok := username.(string)
	if ok {
		fmt.Printf("%v, %T", val, val) // Jack, string
	} else {
		fmt.Println("获取失败")
	}
	ctx.String(http.StatusOK, "this is admin profile")
}
