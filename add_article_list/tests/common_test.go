package tests

import (
	"LearnGin/add_article_list/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []models.Articles

func TestMain(m *testing.M)  { // 在执行test函数之前启动
	// 设置gin为测试模式
	gin.SetMode(gin.TestMode)
	// 执行其它的test
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine{
	r:=gin.Default()
	if withTemplates { // 是否加载模板
		r.LoadHTMLGlob("../templates/*")
	}
	return r
}

// 处理请求并对它的响应做测试
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool){
	w:=httptest.NewRecorder() // 创建一个新的响应体
	r.ServeHTTP(w, req) // 创建和启动服务，处理上面的请求并将响应结果写入w

	if !f(w){ // f是一个函数，返回值表明测试成功或失败 --> 通过函数的调用可以测试各种不同的HTTP请求，减少代码冗余
		t.Fail()
	}
}

// 将初始的文章列表存储到临时文章列表中
func saveLists(){
	tmpArticleList=models.ArticleList
} 

// 在测试结束后用临时文章列表来恢复源文章列表
func restoreLists(){
	models.ArticleList=tmpArticleList
}