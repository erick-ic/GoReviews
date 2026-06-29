package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(ctx *gin.Context) {
	fmt.Println("InitMiddleware...")
}

func Middleware1(ctx *gin.Context) {
	fmt.Println("Middleware1...")
}
func Middleware2(ctx *gin.Context) {
	fmt.Println("Middleware2...")
}
func Middleware3(ctx *gin.Context) {
	fmt.Println("Middleware3...")
}
func MiddlewareOne(ctx *gin.Context) {
	fmt.Println("MiddlewareOne")
}
func MiddlewareTwo(ctx *gin.Context) {
	fmt.Println("MiddlewareTwo")
}
func MiddlewareNext(ctx *gin.Context) {
	fmt.Println("Middleware--1")
	ctx.Next()
	//先调用当前中间件之后的方法，再打印Middleware--2
	fmt.Println("Middleware--2")
}

func MiddlewareAbort(ctx *gin.Context) {
	fmt.Println("Middleware--1")
	ctx.Abort()
	//不调用当前中间件之后的方法
	fmt.Println("Middleware--2")
}

func main() {
	router := gin.Default()
	router.GET("/middleware", InitMiddleware, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is middleware")
		fmt.Println("start...")
	})
	//InitMiddleware...
	//start...

	router.GET("/many/middleware", Middleware1, Middleware2, Middleware3, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is many middleware")
		fmt.Println("start...")
	})
	//Middleware1...
	//Middleware2...
	//Middleware3...
	//start...

	//router.Use(MiddlewareOne, MiddlewareTwo)
	router.GET("/globalMiddleware", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is globalMiddleware")
		fmt.Println("start...")
	})
	//MiddlewareOne
	//MiddlewareTwo
	//start...

	router.GET("/middlewareNext", MiddlewareNext, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is middlewareNext")
		fmt.Println("start...")
	})
	//Middleware--1
	//start...
	//Middleware--2

	router.GET("/middlewareAbort", MiddlewareAbort, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is middlewareAbort")
		fmt.Println("start...")
	})
	//Middleware--1
	//Middleware--2
	router.Run(":8080")
}
