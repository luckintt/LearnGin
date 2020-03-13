package main

import (
	"LearnGin/add_article_list/handlers"
)

func initializeRoutes()  {
	// 主页显示
	router.GET("/", handlers.ShowIndexPage) // 将route handler也单独定义

	// 显示每个文章对应的视图
	router.GET("/article/view/:article_id", handlers.GetArticleByID)
}
