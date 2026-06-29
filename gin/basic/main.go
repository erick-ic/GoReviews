package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Subscribe int    `json:"subscribe"`
}

func main() {

	server := gin.Default()

	//配置路由
	//ping
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	//pong

	//增
	server.POST("/create", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Create sunccess!")
	})
	//Create sunccess!

	//删
	server.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Delete sunccess!")
	})
	//Delete sunccess!

	//改
	server.PUT("/update", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Update sunccess!")
	})
	//Update sunccess!

	//查
	server.GET("/read", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Read sunccess!")
	})
	//Read sunccess!

	//XML
	server.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"code":    200,
			"message": "xml",
		})
	})
	//<map>
	//<code>200</code>
	//<message>xml</message>
	//</map>

	//HTML 模版语法
	//templates/* 只匹配 templates 目录下的一级文件。
	//如果demo.html 放在 templates/demo/index.html，则必须使用 templates/**/* 或 templates/demo/*
	/*生产环境示例：
	你的项目根目录/
	├── client.go
	└── templates/
	├── admin/
	│   └── dashboard.html
	└── user/
	└── profile.html
	*/
	server.LoadHTMLGlob("templates/*")
	server.GET("/html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "demo.html", gin.H{
			"code":    200,
			"message": "hello demo.html",
		})
	})
	//<!DOCTYPE html>
	//<html lang="en">
	//
	//<head>
	//<meta charset="UTF-8">
	//<title>Title</title>
	//</head>
	//
	//<body>
	//<div>this is demo...</div>
	//</body>

	//JSON
	server.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "json",
		})
	})
	//{
	//	"code": 200,
	//	"message": "json"
	//}

	//获取文章
	server.GET("/article", func(ctx *gin.Context) {
		article := &Article{
			Title:     "文章标题",
			Desc:      "文章简介",
			Subscribe: 21,
		}
		ctx.JSON(200, article)
	})
	// {
	// 	"title": "文章标题",
	// 	"desc": "文章简介",
	// 	"subscribe": 21
	// }

	//JSONP
	server.GET("/jsonp", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"code":    200,
			"message": "jsonp",
		}
		ctx.JSONP(http.StatusOK, data)
	})
	//{
	//	"code": 200,
	//	"message": "jsonp"
	//}

	server.Run(":8080")
}
