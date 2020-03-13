package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main()  {
	router=gin.Default()

	router.LoadHTMLGlob("templates/*") // 加载html模板

	// add_article_list\main.go:12:2: undefined: initializeRoutes ： edit configuration -> files -> 添加routes.go文件 ->working directory 改为response_format
	initializeRoutes() // 因为路由会有很多，所以将它拆分到一个新文件中，来处理所有的route

	router.Run()
}
