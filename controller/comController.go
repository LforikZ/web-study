package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-study/entity"
	"web-study/service"
)

func InsertComData(c *gin.Context) {
	//1、处理请求参数
	p := new(entity.ParamInsertCommunity)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		fmt.Println(p)
		zap.L().Error("InsertComData() with invalid param", zap.Error(err))
		//判断err是不是
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errors.Translate(trans))
		return
	}
	//2、业务处理
	if err := service.InsertComData(p); err != nil {
		if errors.Is(err, service.ErrorComIDExit) {
			ResponseError(c, CodeComIDExit)
			return
		}
		if errors.Is(err, service.ErrorComNameExit) {
			ResponseError(c, CodeComNameExit)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, err)
		return
	}
	//3、返回响应
	ResponseSuccess(c, CodeSuccess)
	return
}

func GetCommunityList(c *gin.Context) {
	list := service.GetCommunityList()
	ResponseSuccess(c, list)
	return
}
