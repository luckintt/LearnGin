package tests

import (
	"LearnGin/final_version/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 当用户没有登录时测试中间件EnsureLoggedIn()
func TestEnsureLoggedInUnauthenticated(t *testing.T){
	r:=getRouter(false) // 不加载模板
	r.GET("/", setLoggedIn(false), middleware.EnsureLoggedIn(), func(c *gin.Context) {
		/*
		* 调用中间件setLoggedIn将is_logged_in置为false, 因为用户没有登录，所以不会执行该handler
		* 若执行了，则证明中间件EnsureLoggedIn出错了(EnsureLoggedIn是确保用户登录了，若没有登录则会未认证)
		*/
		t.Fail()
	})
	// 使用helper方法执行请求并对响应结果进行测试
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

// 当用户登录时测试中间件EnsureLoggedIn()
func TestEnsureLoggedInAuthenticated(t *testing.T){
	r:=getRouter(false)
	r.GET("/", setLoggedIn(true), middleware.EnsureLoggedIn(), func(c *gin.Context) {
		// 调用中间件setLoggedIn将is_logged_in置为true, 则用户登录成功，会执行该handler
		c.Status(http.StatusOK)
	})

	testMiddlewareRequest(t, r, http.StatusOK)
}

// 当用户登录时测试中间件EnsureNotLoggedIn()
func TestEnsureNotLoggedInAuthenticated(t *testing.T){
	r:=getRouter(false)
	r.GET("/", setLoggedIn(true), middleware.EnsureNotLoggedIn(), func(c *gin.Context) {
		// 调用中间件setLoggedIn将is_logged_in置为true, 用户登录成功
		// 中间件EnsureNotLoggedIn要保证用户没有登录，因此不会执行该handler
		t.Fail()
	})
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

// 当用户没有登录时测试中间件EnsureNotLoggedIn()
func TestEnsureNotLoggedInUnauthenticated(t *testing.T){
	r:=getRouter(false)
	r.GET("/", setLoggedIn(false), middleware.EnsureNotLoggedIn(), func(c *gin.Context) {
		// 用户没有登录，满足中间件EnsureNotLoggedIn的要求，因此会执行该handler
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

// 这是一个用于设置"is_logged_in"参数的中间件，仅仅用于测试
func setLoggedIn(b bool) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}
