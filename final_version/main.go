package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main(){
	// 设置为release模式，加快运行速度
	gin.SetMode(gin.ReleaseMode)

	// 这里如果用:=会报错，是因为
	// 对于使用:=定义的变量，如果新变量router与那个同名已定义变量 (这里就是那个全局变量router)不在一个作用域中时
	// 那么golang会新定义这个变量router，遮盖住全局变量router，这就是导致nil pointer问题的真凶。
	router=gin.Default()

	router.LoadHTMLGlob("templates/*") // 加载模板

	initializeRoutes()

	router.Run()
}