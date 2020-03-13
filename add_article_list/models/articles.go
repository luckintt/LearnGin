package models

import "errors"

type Articles struct {
	ID		int  	`json:"id"`
	Title	string  `json:"title"`
	Content string  `json:"content"`
}

// 必须用var来定义，var定义的是全局变量，:=定义的是局部变量，在getAllArticles()中只能调用外面的全局变量
var ArticleList = []Articles{
	Articles{ID:1, Title:"Article 1", Content:"Article 1 body"},
	Articles{ID:2, Title:"Article 2", Content:"Article 2 body"},
} // 这个数据应该是从数据库或者静态文件中读取，这里为了简单，直接定义为变量

func GetAllArticles() []Articles{
	return ArticleList
}

func GetArticleByID(id int) (*Articles, error){
	for _, v := range ArticleList{
		if v.ID==id {
			return &v, nil
		}
	}
	return nil, errors.New("Article Not Found")
}
