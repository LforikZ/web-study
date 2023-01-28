package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web-study/controller"
	"web-study/entity"
	"web-study/service"
)

// InsertPostData
// @Description 上传帖子
// @Author Zihao_Li 2023-01-28 14:10:22
// @Param c
func InsertPostData(c *gin.Context) {
	//获取参数校验
	p := new(entity.ParamPostData)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("InsertPostData() with invalid param", zap.Error(err))
		//判断err是不是
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, errors.Translate(controller.Trans))
		return
	}
	//从c中取出id
	userID, err := controller.GetCurrentUser(c)
	if err != nil {
		zap.L().Debug(" controller.GetCurrentUser(c) fieled", zap.Any("err", err))
		controller.ResponseError(c, controller.CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//创建帖子
	if err := service.CreatPost(p); err != nil {
		fmt.Println(err)
		zap.L().Error("CreatPost() detail", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	//返回响应
	controller.ResponseSuccess(c, nil)
	return
}

// GetPostList
// @Description 获取多个帖子
// @Author Zihao_Li 2023-01-28 14:11:03
// @Param c
func GetPostList(c *gin.Context) {
	//获取分页参数
	page, size := controller.GetPageInfo(c)
	//获取数据
	list, err := service.GetPostList(page, size)
	if err != nil {
		fmt.Println(err)
		controller.ResponseError(c, controller.CodeGetListFiled)
		return
	}
	controller.ResponseSuccess(c, list)
	return
}

// GetPostData
// @Description  获取单个帖子内容
// @Author Zihao_Li 2023-01-28 14:11:14
// @Param c
func GetPostData(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		zap.L().Error("get post data detail", zap.Error(err))
		controller.ResponseError(c, controller.CodeInvalidParam)
		return
	}
	if id == 0 {
		controller.ResponseError(c, controller.CodePostNotExit)
		return
	}
	data, err := service.GetPostDataById(id)
	if err != nil {
		zap.L().Error("service.GetPostData(id) failed", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, data)
	return
}
