package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件：确保用户若已经登录则触发异常，只有未登录的用户才能使用
func EnsureNotLoggedIn() gin.HandlerFunc{
	return func(c *gin.Context) {
		loggedInInterface, _:=c.Get("is_logged_in")
		// 左边可以是一个值，这个时候如果接口保管的不是该类型会panic;也可以是两个值，此时如果接口保管的类型精确为括号内的类型第二个值为真，否则为假
		loggedIn:=loggedInInterface.(bool) // 判断该接口是不是bool类型，不是则panic
		if loggedIn{ // is_logged_in=true
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureLoggedIn()gin.HandlerFunc  {
	return func(c *gin.Context) {
		loggedInInterface,_:=c.Get("is_logged_in")
		loggedIn:=loggedInInterface.(bool)
		if !loggedIn{ // is_logged_in=false
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// 用中间件设置用户是否登录
func SetUserStatus() gin.HandlerFunc { // 中间件的返回值是gin.HandleFunc
	return func(c *gin.Context){
		if token, err:=c.Cookie("token");err==nil && token!=""{
			c.Set("is_logged_in", true)
		}else{
			c.Set("is_logged_in", false)
		}
	}
}
