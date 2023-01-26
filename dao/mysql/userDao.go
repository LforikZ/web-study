package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"web-study/entity"
)

type User struct {
	gorm.Model
	UserID   int64  `json:"userId" gorm:"user_id"`
	Gender   int    `json:"gender" gorm:"gender"`
	UserName string `json:"userName" gorm:"user_name"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
}

// SelectByUsername 通过username 来查询用户是否存在
func SelectByUsername(username string) User {
	var user User
	db.Where("user_name=?", username).Find(&user)
	return user
}

// InsertUser 向数据库插入一条新的用户记录
func InsertUser(user *entity.User) (err error) {
	//加密
	password := Md5Password(user.UserName, user.Password)
	user.Password = password
	//插入
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}
	return
}

// Md5Password md5加密算法
func Md5Password(username string, password string) string {
	h := md5.New()
	h.Write([]byte(username))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
