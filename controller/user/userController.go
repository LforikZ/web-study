package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web-study/controller"
	"web-study/entity"
	"web-study/service"
)

// SignUpUser
// @Description 注册用户
// @Author Zihao_Li 2023-01-28 15:14:23
// @Param c
func SignUpUser(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(entity.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("SignUpUser() with invalid param", zap.Error(err))
		//判断err是不是
		er, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, er.Translate(controller.Trans))
		return
	}

	//2.业务处理
	err := service.SignUpUser(p)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, service.ErrorUserExit) {
			controller.ResponseError(c, controller.CodeUserExit)
		}
		return
	}
	//3.返回响应
	controller.ResponseSuccess(c, controller.CodeSuccess)
	return
}

// LoginUp
// @Description 登录用户
// @Author Zihao_Li 2023-01-28 15:14:33
// @Param c
func LoginUp(c *gin.Context) {
	//获取参数校验
	p := new(entity.ParamLoginUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("LoginUp() with invalid param", zap.Error(err))
		//判断err是不是
		er, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, er.Translate(controller.Trans))
		return
	}
	//业务处理
	apiuser, err := service.Login(p)
	if err != nil {
		zap.L().Error("login失败", zap.Error(err))
		if errors.Is(err, service.ErrorUserNotExit) {
			controller.ResponseError(c, controller.CodeUserNotExit)
		} else if errors.Is(err, service.ErrorInvalidPassword) {
			controller.ResponseError(c, controller.CodeInvalidPassword)

		}
		return
	}
	controller.ResponseSuccess(c, gin.H{
		"user_id":   strconv.FormatInt(apiuser.UserID, 10), //前端js最大识别为 2的53-1次方 int64最大为 2的63-1; 如果不处理可能会阈值
		"user_name": apiuser.UserName,
		"token":     apiuser.Token,
	})
	return
}
