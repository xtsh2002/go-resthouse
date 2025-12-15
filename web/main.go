package main

// 导入必要的包
import (
	"net/http"
	"restproject/web/controller"

	"github.com/gin-gonic/gin"
)

// 添加gin框架开发3步骤
func main() {
	// 创建一个新的Gin引擎
	r := gin.Default()

	// 定义一个路由，当访问根路径时返回"Hello World!"
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "租房项目开始了!")
	})
	r.Static("/home", "view")
	r.GET("/api/v1.0/session", controller.GetSeession)
	r.GET("/api/v1.0/imagecode/:uuid", controller.GetCaptcha)
	// 让服务器在0.0.0.0:8080上运行并监听和响应请求
	r.Run(":8080") // 注意这里的端口号要和上面的代码保持一致
}
