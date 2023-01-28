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

// InsertComData
// @Description 创建社区
// @Author Zihao_Li 2023-01-28 15:52:44
// @Param c
func InsertComData(c *gin.Context) {
	//1、处理请求参数
	p := new(entity.ParamCommunity)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		fmt.Println(p)
		zap.L().Error("InsertComData() with invalid param", zap.Error(err))
		//判断err是不是
		er, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, er.Translate(controller.Trans))
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

// GetCommunityList
// @Description 获取所有社区
// @Author Zihao_Li 2023-01-28 15:52:57
// @Param c
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

// GetCommunityById
// @Description 根据id获取社区内容
// @Author Zihao_Li 2023-01-28 15:53:20
// @Param c
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
