package main

import (
	"github.com/gin-contrib/sessions"       // 提供会话（Session）管理接口
	"github.com/gin-contrib/sessions/redis" // 基于 Redis 的会话存储引擎实现
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// 1. 创建基于 Redis 的会话存储引擎
	//    参数说明：
	//    - 16       : 连接池中最大空闲连接数（Redis 默认支持 16 个数据库，此处指连接数）
	//    - "tcp"    : 网络协议
	//    - "localhost:6379" : Redis 服务地址
	//    - ""       : 用户名（Redis 6.0+ 支持 ACL，空表示无需用户名）
	//    - ""       : 密码（空表示无密码）
	//    - []byte("3akQBTZmfkuEjQacH5hvUynDnmPvAf7Y") : 身份验证密钥（用于签名 Session ID，防篡改）
	//    - []byte("Z4d8tz8WDKXT3AvYJkmhEb5VEFfxHHS2") : 数据加密密钥（用于加密会话数据，保证数据私密性）
	store, _ := redis.NewStore(
		16,
		"tcp",
		"localhost:6379",
		"",
		"",
		[]byte("3akQBTZmfkuEjQacH5hvUynDnmPvAf7Y"),
		[]byte("Z4d8tz8WDKXT3AvYJkmhEb5VEFfxHHS2"))

	// 2. 将会话中间件注册到 Gin 引擎
	//    sessions.Sessions("mysession", store) 创建一个中间件，指定 Cookie 名称为 "mysession"
	//    中间件会在每个请求前自动从 Cookie 中读取会话 ID，并到 Redis 中加载会话数据；
	//    若 Cookie 无效或不存在，则创建一个新的空会话。
	server.Use(sessions.Sessions("mysession", store))

	server.GET("/hello", func(ctx *gin.Context) {
		// 实际开发中，以下代码通常放在登录成功后执行

		// 3. 从请求上下文中获取当前会话对象
		session := sessions.Default(ctx)

		// 4. 在会话中存储键值对：userId = "1"
		session.Set("userId", "1")

		// 5. 设置当前会话关联的 Cookie 属性
		//     MaxAge: 10 表示 Cookie 有效期为 10 秒（单位：秒）
		//     浏览器将在 10 秒后自动删除该 Cookie
		session.Options(sessions.Options{
			MaxAge: 10,
		})

		// 6. 保存会话数据至 Redis，并生成 Set-Cookie 响应头
		//     若未调用 Save()，则数据不会持久化，客户端也收不到 Cookie
		session.Save()

		ctx.JSON(200, gin.H{
			"userId": "1",
			"code":   200,
			"msg":    "ok",
		})
	})

	server.Run(":8080")
}
