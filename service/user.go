package service

import (
	"errors"
	"web-study/dao/mysql"
	"web-study/entity"
	"web-study/pkg/snowflake"
)

var (
	ErrorUserExit        = errors.New("用户已存在")
	ErrorUserNotExit     = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func SignUpUser(p *entity.ParamSignUp) (err error) {
	//判断是否存在
	result := mysql.SelectByUsername(p.Username)
	if result.UserID != 0 {
		return ErrorUserExit
	}
	//生成UID
	userId := snowflake.GenID()
	//构建一个User实例
	user := entity.User{
		UserID:   userId,
		UserName: p.Username,
		Password: p.Password,
	}
	//保存到数据库
	return mysql.InsertUser(&user)
}

func Login(p *entity.ParamLoginUp) (err error) {
	result := mysql.SelectByUsername(p.Username)
	if result.UserID == 0 {
		return ErrorUserNotExit
	}
	password := mysql.Md5Password(p.Username, p.Password)
	if password != result.Password {
		return ErrorInvalidPassword
	}
	return
}
