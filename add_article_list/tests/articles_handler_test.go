package tests

import (
	"LearnGin/add_article_list/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T)  {
	r:=getRouter(true) // 有模板

	r.GET("/", handlers.ShowIndexPage) // 使用相同的handler

	req, _:=http.NewRequest("GET","/", nil) // 用该路由创建一个请求体

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool{ // 调用testHTTPResponse时已经将返回的结果写入w
		statusOK:=(w.Code==http.StatusOK)  // 测试http请求的状态码是不是200
		p,err:=ioutil.ReadAll(w.Body) // 读取响应体中的全部数据
		pageOK:=(err==nil && (strings.Index(string(p), "<title>Home Page</title>")>0)) // 测试页面标题是Home Page
		return statusOK && pageOK
	})
}

func TestGetArticleByIDUnauthenticated(t *testing.T)  {
	r:=getRouter(true)

	r.GET("/article/view/:article_id", handlers.GetArticleByID)

	req, _:=http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool{
		statusOK:=(w.Code==http.StatusOK)
		rep,err:=ioutil.ReadAll(w.Body)
		pageOk:=(err==nil && strings.Index(string(rep), "<title>Article 1</title>")>0)
		return statusOK && pageOk
	})
}
