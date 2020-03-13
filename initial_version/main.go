package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine
func main(){
	router=gin.Default() // 使用gin默认的路由

	router.LoadHTMLGlob("templates/*")  // 加载templates文件夹下的所有模板

	router.GET("/", func(c *gin.Context){ // 注册路由，使用get方式请求
		c.HTML( // 调用上下文的HTML方法来渲染模板
			http.StatusOK,  // 状态码
			"index.html", // 模板名
			gin.H{ // 将数据传到页面
				"title":"Home Page",
			},
		)
	})

	router.Run() // 默认是localhost:8080
}
