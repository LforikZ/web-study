package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-study/controller"
	"web-study/entity"
	"web-study/service"
)

// SignUpUser
/**
 * @Author: ZiHao-Li
 * @Description: 注册用户
 * @Date: 2023/1/5 23:45
 * @Param:
 * @return:
 **/
func SignUpUser(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(entity.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("SignUpUser() with invalid param", zap.Error(err))
		//判断err是不是
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, errors.Translate(controller.Trans))
		return
	}

	//2.业务处理
	err := service.SignUpUser(p)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, service.ErrorUserExit) {
			controller.ResponseErrorWithMsg(c, controller.CodeUserExit, "注册失败")
		}
		return
	}
	//3.返回响应
	controller.ResponseSuccess(c, controller.CodeSuccess)
	return
}

// LoginUp
/**
 * @Author: ZiHao-Li
 * @Description:  登录功能
 * @Date: 2023/1/15 18:12
 * @Param:
 * @return:
 **/
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
	token, err := service.Login(p)
	if err != nil {
		zap.L().Error("login失败", zap.Error(err))
		if errors.Is(err, service.ErrorUserNotExit) {
			controller.ResponseError(c, controller.CodeUserNotExit)
		} else if errors.Is(err, service.ErrorInvalidPassword) {
			controller.ResponseError(c, controller.CodeInvalidPassword)
		}
		return
	}
	controller.ResponseSuccess(c, token)
	return
}