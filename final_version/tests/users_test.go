package tests

import (
	"LearnGin/final_version/models"
	"testing"
)

// 测试不同的用户名密码组合
func TestUserValidity(t *testing.T) {
	if !models.IsValid("user1", "pass1") {
		t.Fail()
	}

	if models.IsValid("user2", "pass1") {
		t.Fail()
	}

	if models.IsValid("user1", "") {
		t.Fail()
	}

	if models.IsValid("", "pass1") {
		t.Fail()
	}

	if models.IsValid("User1", "pass1") {
		t.Fail()
	}
}

// 测试有效的用户名密码注册
func TestValidUserRegistration(t *testing.T)  {
	saveLists()

	u, err:=models.RegisterNewUser("newname", "newpass")
	if err!=nil || u.Username==""{
		t.Fail()
	}

	restoreLists()
}

// 测试无效的用户名密码注册
func TestInvalidUserRegistration(t *testing.T){
	saveLists()

	u, err:=models.RegisterNewUser("user1", "pass1") // 正确结果是:err!=nil && u=nil
	if err==nil  || u!=nil{
		t.Fail()
	}

	u, err=models.RegisterNewUser("newname", "")
	if err==nil || u!=nil{
		t.Fail()
	}

	restoreLists()
}

func TestUsernameAvailability(t *testing.T) {
	saveLists()

	if models.IsUsernameAvailable("user1"){
		t.Fail()
	}

	if !models.IsUsernameAvailable("newname"){
		t.Fail()
	}

	models.RegisterNewUser("newuser", "newpass")
	if models.IsUsernameAvailable("newuser"){
		t.Fail()
	}

	restoreLists()
}

