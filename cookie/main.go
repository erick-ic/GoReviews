package main

import (
	"github.com/gin-contrib/sessions"        // 提供会话（Session）管理接口
	"github.com/gin-contrib/sessions/cookie" // 导入 cookie 驱动包
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// 1. 创建一个基于客户端 Cookie 的会话存储引擎
	//    - cookie.NewStore 来自 github.com/gin-contrib/sessions/cookie
	//    - 参数 []byte("secret") 是用于签名 Session ID 的密钥（防止篡改），生产环境必须使用复杂且保密的字符串
	store := cookie.NewStore([]byte("secret"))
	// 2. 将会话中间件注册到 Gin 引擎，作用于所有请求
	//    - sessions.Sessions("mysession", store) 创建一个中间件，指定 Cookie 名称为 "mysession"
	//    - server.Use(...) 使该中间件在每次请求时自动执行，负责从请求头读取/解密 Cookie，并将会话对象注入上下文
	server.Use(sessions.Sessions("mysession", store))

	server.GET("/hello", func(ctx *gin.Context) {
		//实际开发中，以下代码通常放在登录成功后执行

		// 3. 从当前请求的上下文中获取默认的会话对象
		//    - sessions.Default(ctx) 会检查请求中是否有有效的 "mysession" Cookie，若有则还原会话数据；否则新建一个空会话
		//    - 返回的 sess 是一个会话对象，提供了 Get/Set/Delete/Options/Save 等方法
		session := sessions.Default(ctx)
		// 4. 在会话中存储键值对数据
		//    - 键 "userId"，值是从登录验证服务获得的用户唯一标识（u.Id）
		//    - 数据暂存于内存，尚未持久化，需调用 Save() 才会写入 Cookie
		session.Set("userId", "1")
		// 5. 设置当前会话关联的 Cookie 属性（有效期等）
		//    - sessions.Options 结构体用于控制 Cookie 的 Path/Domain/MaxAge/Secure/HttpOnly/SameSite 等
		//    - MaxAge: 10 表示 Cookie 有效期为 10 秒（单位：秒），浏览器会在 10 秒后自动删除该 Cookie
		//    - 若不设置 MaxAge，默认值为 0，表示会话级 Cookie（浏览器关闭即失效）
		session.Options(sessions.Options{
			MaxAge: 10, //单位秒
		})
		// 6. 保存会话数据，真正触发持久化操作
		//    - 对于 cookie.NewStore，会将会话数据序列化、加密并签名，然后生成 Set-Cookie 响应头发送给客户端
		//    - 必须显式调用 Save()，否则数据不会实际写入 Cookie，客户端无法收到任何会话凭证
		session.Save()

		ctx.JSON(200, gin.H{
			"userId": "1",
			"code":   200,
			"msg":    "ok",
		})
	})

	server.Run("localhost:8080")
}
