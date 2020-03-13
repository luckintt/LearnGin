package models

import (
	"errors"
	"strings"
)

type Users struct {
	Username 	string 		`json:"username"`
	Password	string 		`json:"password"`
}

var UserList=[]Users{
	{Username:"user1", Password:"pass1"},
	{Username:"user2", Password:"pass2"},
	{Username:"user3", Password:"pass3"},
}

// 判断用户名密码是否正确
func IsValid(username string, password string) bool {
	for _, v:=range UserList{
		if v.Username==username && v.Password==password{
			return true
		}
	}
	return false
}

// 用给定的用户名密码注册一个新用户
func RegisterNewUser(username string, password string) (*Users, error) {
	if strings.TrimSpace(password)==""{
		return nil, errors.New("The password can't be empty")
	}else if !IsUsernameAvailable(username){
		return nil, errors.New("The username isn't avaliable")
	}

	u:=Users{Username:username, Password:password}
	UserList=append(UserList, u)
	return &u, nil
}

// 检测用户名是否可用
func IsUsernameAvailable(username string) bool {
	for _,v:=range UserList{
		if v.Username==username{
			return false
		}
	}
	return true
}