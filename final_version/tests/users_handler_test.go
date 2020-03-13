package tests

import (
	"LearnGin/final_version/handlers"
	"LearnGin/final_version/middleware"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// 对于已经登录的用户测试ShowLoginPage函数
func TestShowLoginPageAuthenticated(t *testing.T)  {
	// 创建一个响应体
	w:=httptest.NewRecorder()
	// 设置token模拟一个认证的用户
	http.SetCookie(w, &http.Cookie{Name:"token", Value:"123"})
	// 创建一个路由
	r:=getRouter(true)
	// 使用中间件(这里用了两个中间件，若只使用第二个中间件会报错)和router handler func
	r.GET("/user/login", middleware.SetUserStatus(),middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)
	// 对这个路由发起请求
	req, _:=http.NewRequest("GET","/user/login", nil)
	req.Header=http.Header{"Cookie":w.HeaderMap["Set-Cookie"]}

	// 启动一个服务并且响应上面的请求
	r.ServeHTTP(w, req)
	// 因为用户已经登录了，再登录应该返回unauthorized
	if w.Code!=http.StatusUnauthorized{
		t.Fail()
	}
}

// 对于没有登录的用户测试ShowLoginPage函数
func TestShowLoginPageUnauthenticated(t *testing.T){
	// 创建路由
	r:=getRouter(true)
	r.GET("/user/login", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)

	req, _:= http.NewRequest("GET", "/user/login", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK:=w.Code==http.StatusOK
		body, err:=ioutil.ReadAll(w.Body)
		return  statusOK && err==nil && strings.Index(string(body), "<title>Login Page</title>")>0
	})
}

// 对于登录的用户测试PerformLogin函数
func TestPerformLoginAuthenticated(t *testing.T){
	w:=httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name:"token", Value:"123"})
	r:=getRouter(true)
	r.POST("/user/login", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.PerformLogin)
	// 登录传递的参数
	loginPayLoad:=getLoginPOSTPayload() // 该函数的用户名密码正确
	req, _:=http.NewRequest("POST","/user/login", strings.NewReader(loginPayLoad))
	req.Header=http.Header{"Cookie":w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayLoad)))
	r.ServeHTTP(w, req)
	// 因为用户已经登录，不满足中间件EnsureLoggedIn，所以最后会返回未认证
	if w.Code!=http.StatusUnauthorized{
		t.Fail()
	}
}

// 对于没有登录的用户测试PerformLogin函数
func TestPerformLoginUnauthenticated(t *testing.T){
	r:=getRouter(true)
	r.POST("/user/login", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

	loginPayLoad:=getLoginPOSTPayload()
	req, _:=http.NewRequest("POST", "/user/login", strings.NewReader(loginPayLoad))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")  // 要有这两个才能成功的发送POST请求
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayLoad)))
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK:=w.Code==http.StatusOK
		body, err:=ioutil.ReadAll(w.Body)
		return statusOK && err==nil && strings.Index(string(body), "<title>Login Successful</title>")>0
	})
}

// 对于没有登录的用户，用不正确的用户名密码测试PerformLogin函数
func TestPerformLoginUnauthenticatedIncorrectCredentials(t *testing.T){
	r:=getRouter(true)
	r.POST("/user/login", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

	loginPayLoad:=getRegisterPOSTPayLoad()
	req, _:=http.NewRequest("POST", "/user/login", strings.NewReader(loginPayLoad))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")  // 要有这两个才能成功的发送POST请求
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayLoad)))
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 用户名密码不正确，不会登录成功，而且会返回StatusBadRequest
		return w.Code==http.StatusBadRequest
	})
}

// 对于已经登录的用户，测试ShowRegisterPage函数
func TestShowRegisterPageAuthenticated(t *testing.T){
	w:=httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name:"token", Value:"123"})

	r:=getRouter(true)
	r.GET("/user/register", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.ShowRegisterPage)

	req, _:=http.NewRequest("GET", "/user/register", nil)
	req.Header=http.Header{"Cookie":w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)
	if w.Code!=http.StatusUnauthorized { // 已经登录的用户不能查看注册页面，在执行中间件EnsureNotLoggedIn时会返回未认证
		t.Fail()
	}
}

// 对于没有登录的用户，测试ShowRegisterPage函数
func TestShowRegisterPageUnauthenticated(t *testing.T){
	r:=getRouter(true)
	r.GET("/user/register", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.ShowRegisterPage)

	registerPayLoad:=getRegisterPOSTPayLoad()
	req,_:= http.NewRequest("GET", "/user/register", strings.NewReader(registerPayLoad))

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK:=w.Code==http.StatusOK
		body, err:=ioutil.ReadAll(w.Body)
		return statusOK && err==nil && strings.Index(string(body), "<title>Register Page</title>")>0
	})
}

// 对于已经登录的用户，测试PerformRegister函数
func TestPerformRegisterAuthenticated(t *testing.T){
	w:=httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name:"token", Value:"123"})

	r:=getRouter(true)
	r.POST("/user/register", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.PerformRegister)

	registerPayLoad:=getRegisterPOSTPayLoad()
	req, _:=http.NewRequest("POST", "/user/register", strings.NewReader(registerPayLoad))
	req.Header=http.Header{"Cookie":w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-wwww-form-urlencoded")  // 要有这两个才能成功的发送POST请求
	req.Header.Add("Content-Length", strconv.Itoa(len(registerPayLoad)))

	//  已经登录的用户不能进行，在执行中间件EnsureNotLoggedIn时会返回未认证
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code==http.StatusUnauthorized
	})
}

// 对于没有登录的用户，测试PerformRegister函数
func TestPerformRegisterUnauthenticated(t *testing.T){
	r:=getRouter(true)
	r.POST("/user/register", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.PerformRegister)

	registerPayLoad:=getRegisterPOSTPayLoad()
	req, _:=http.NewRequest("POST", "/user/register", strings.NewReader(registerPayLoad))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // 要有这两个才能成功的发送POST请求
	req.Header.Add("Content-Length", strconv.Itoa(len(registerPayLoad)))

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK:=w.Code==http.StatusOK
		body, err:=ioutil.ReadAll(w.Body)
		return statusOK && err==nil && strings.Index(string(body), "<title>Register &amp;&amp; Login Successful</title>")>0
	})
}

// 用无效的用户名来进行注册
func TestPerformRegisterUnauthenticatedUnavailableUsername(t *testing.T){
	r:=getRouter(true)
	r.POST("/user/register", middleware.SetUserStatus(), middleware.EnsureNotLoggedIn(), handlers.PerformRegister)

	registerPayLoad:=getLoginPOSTPayload()
	req, _:=http.NewRequest("POST", "/user/register", strings.NewReader(registerPayLoad))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registerPayLoad)))

	// 因为用户名无效，所以在执行handler时会返回StatusBadRequest
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code==http.StatusBadRequest
	})
}

func getLoginPOSTPayload() string {
	param:=url.Values{}
	param.Add("username", "user1")
	param.Add("password","pass1")
	return param.Encode()
}

func getRegisterPOSTPayLoad() string {
	param:=url.Values{}
	param.Add("username", "u1")
	param.Add("password","p1")
	return param.Encode()
}
