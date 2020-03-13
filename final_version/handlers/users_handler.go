package handlers

import (
	"LearnGin/final_version/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func ShowLoginPage(c *gin.Context)  {
	Render(c, gin.H{"title":"Login Page"}, "login.html")
}

func PerformLogin(c *gin.Context)  {
	username:=c.PostForm("username")
	password:=c.PostForm("password")

	valid:=models.IsValid(username, password)
	var sameSiteCookie http.SameSite
	if valid{
		// 在cookie中设置token
		token:=GenerateSessionToken()
		// maxAge是cookie的有效时间
		c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		c.Set("is_logged_in", true)

		Render(c, gin.H{"title":"Login Successful"}, "login_successful.html")
	}else{
		// 如果用户名密码错误，就在登录页面显示错误信息
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"ErrorTitle":"Login Failed", "ErrorMessage":"Invalid credentials provided"})
	}
}

func GenerateSessionToken() string {
	// 随机生成的为Int63的数字，然后将其转换成string作为会话token
	// 使用这种方式生成会话token是不安全的，不建议在开发中使用
	return strconv.FormatInt(rand.Int63(),16)
}

func Logout(c *gin.Context)  {
	var sameSiteCookie http.SameSite
	// 将cookie设置为过期
	c.SetCookie("token", "", -1, "", "", sameSiteCookie, false, true)
	// 重定向到主页
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ShowRegisterPage(c *gin.Context)  {
	Render(c, gin.H{"title":"Register Page"}, "register.html")
}

func PerformRegister(c *gin.Context)  {
	username:=c.PostForm("username")
	password:=c.PostForm("password")

	if _, err:=models.RegisterNewUser(username, password);err==nil{
		token:=GenerateSessionToken()
		var sameSiteCookie http.SameSite
		c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		c.Set("is_logged_in", true)
		Render(c, gin.H{"title":"Register && Login Successful"}, "login_successful.html")
	}else{
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"ErrorTitle":"Register Failed", "ErrorMessage":err.Error()})
	}
}