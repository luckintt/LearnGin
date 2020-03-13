package main

import (
	"LearnGin/final_version/handlers"
	"LearnGin/final_version/middleware"
)

func initializeRoutes() {
	router.Use(middleware.SetUserStatus()) // 对每一个请求都使用中间件(要调用函数),设置用户状态

	router.GET("/", handlers.ShowIndexPage)

	articeRoutes:=router.Group("/article") // 使用花括号{}将分组下的路由包起来，只是为了更加直观，并非必要的。
	{
		//  url: /article/view/some_article_id 处理该url的GET请求
		articeRoutes.GET("/view/:article_id", handlers.GetArticleByID)

		// url: /article/create 处理该url的GET请求
		// 显示新建文章的页面
		// 使用中间件确保用户已经登录
		articeRoutes.GET("/create", middleware.EnsureLoggedIn(), handlers.ShowArticleCreationPage)

		// url: /article/create 处理该url的POST请求
		// 使用中间件确保用户已经登录
		articeRoutes.POST("/create", middleware.EnsureLoggedIn(), handlers.CreateArticle)
	}

	userRoutes:=router.Group("/user")
	{
		// url: /u/login 处理该url的GET请求
		// 显示登录页面
		// 使用中间件确保用户没有登录
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		// url: /u/login 处理该url的POST请求
		// 使用中间件确保用户没有登录
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

		// url: /u/logout 处理该url的GET请求
		// 使用中间件确保用户已经登录
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.Logout)

		// url: /u/register 处理该url的GET请求
		// 显示注册页面
		// 使用中间件确保用户没有登录
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegisterPage)

		// url: /u/register 处理该url的POST请求
		// 使用中间件确保用户没有登录
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.PerformRegister)
	}
}
