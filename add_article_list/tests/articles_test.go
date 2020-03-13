package tests

import (
	"LearnGin/add_article_list/models"
	"testing"
)

// 测试函数要以Test开头再加上函数名，文件名要以_test结尾
func TestGetAllArticles(t *testing.T){ // 单元测试
	alist := models.GetAllArticles() // 若是小写开头就不能调用这个函数
	if(len(alist)!=len(models.ArticleList)){
		t.Fail()
	}
	for i,v := range alist{
		if v.ID!=models.ArticleList[i].ID || v.Title!=models.ArticleList[i].Title || v.Content!=models.ArticleList[i].Content {
			t.Fail()
			break
		}
	}
}

func TestGetArticleByID(t *testing.T)  {
	article, err:=models.GetArticleByID(1)

	if err!=nil || article.ID!=1 || article.Title!="Article 1" || article.Content!="Article 1 body" {
		t.Fail()
	}
}