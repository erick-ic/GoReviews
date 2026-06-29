package routerGroup

import (
	"GoReviews/gin/middlewares"
	"GoReviews/gin/routerController/admin"

	"github.com/gin-gonic/gin"
)

func HandleRouterGroup(router *gin.Engine) {
	//admin
	//adminRouters := router.Group("/admin")
	//{
	//	adminRouters.GET("/index", func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, "this is admin index")
	//	})
	//	adminRouters.GET("/list", func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, "this is admin list")
	//	})
	//}

	//adminRouters := router.Group("/admin")
	//{
	//	adminRouters.GET("/index", admin.UserIndex)
	//	adminRouters.GET("/list", admin.UserList)
	//}

	//adminRouters := router.Group("/admin")
	//{
	//	adminRouters.GET("/index", (&admin.UserController{}).Index)
	//	adminRouters.GET("/list", (&admin.UserController{}).List)
	//}

	adminRouters := router.Group("/admin", middlewares.SetMiddleware)
	{
		adminRouters.GET("/profile", (&admin.ProfileController{}).GetProfile)
	}

	router.Run("127.0.0.1:8080")
}
