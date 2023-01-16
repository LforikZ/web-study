package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"web-study/middlewares"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前用户id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
