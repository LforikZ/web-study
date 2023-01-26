package community

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
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, errors.Translate(controller.Trans))
		return
	}
	//2、业务处理
	if err := service.InsertComData(p); err != nil {
		if errors.Is(err, service.ErrorComIDExit) {
			controller.ResponseError(c, controller.CodeComIDExit)
			return
		}
		if errors.Is(err, service.ErrorComNameExit) {
			controller.ResponseError(c, controller.CodeComNameExit)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, err)
		return
	}
	//3、返回响应
	controller.ResponseSuccess(c, controller.CodeSuccess)
	return
}

func GetCommunityList(c *gin.Context) {
	list, err := service.GetCommunityList()
	if err != nil {
		fmt.Println(err)
		controller.ResponseError(c, controller.CodeGetListFiled)
		return
	}
	controller.ResponseSuccess(c, list)
	return
}

func GetCommunityById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		controller.ResponseError(c, controller.CodeInvalidParam)
		return
	}
	community, err := service.GetCommunityById(id)
	if err != nil {
		controller.ResponseError(c, controller.CodeGetComFiled)
		return
	}
	controller.ResponseSuccess(c, community)
	return
}
