package handlers

import (
	"LearnGin/add_article_list/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowIndexPage(c *gin.Context)  { // 处理器要有一个上下文参数来获取对应的相应信息
	articleList:=models.GetAllArticles()
	// 调用上下文的HTML方法来渲染模板
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{ // 传递参数
			"title":"Home Page",
			"payload":articleList,
		},
	)
}

func GetArticleByID(c *gin.Context){
	//  Param的参数是url的变量名
	if articleID, err:=strconv.Atoi(c.Param("article_id"));err==nil{ // 前面是初始化，后面是if判断
		if article, err:=models.GetArticleByID(articleID); err==nil{
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":article.Title,
					"payload":article,
				},
			)
		}else{
			//  错误请求处理返回要使用c.Abort，不要只是return
			// Abort函数的本质是提前结束后续的handler链条(通过handler的下标索引直接变化为math.MaxInt8/2)但是前面已经执行过的handler链条(包括middleware等)还会继续返回
			c.AbortWithError(http.StatusNotFound, err)
		}
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}
}
