// @Author Zihao_Li 2023/1/29 16:49:00
package vote

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-study/controller"
	"web-study/entity"
	"web-study/service"
)

// VoteDataChange 投票帖子接口
// @Summary 投票帖子接口
// @Description 可以投喜欢或者讨厌票
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query entity.ParamVoteData true "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseVote
// @Router /vote/change [post]
func VoteDataChange(c *gin.Context) {
	// 参数校验
	p := new(entity.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			controller.ResponseError(c, controller.CodeInvalidParam)
			return
		}
		errData := controller.RemoveTopStruct(errs.Translate(controller.Trans)) // 翻译并去除错误中结构体名字
		controller.ResponseErrorWithMsg(c, controller.CodeInvalidParam, errData)
		return
	}

	//业务处理
	//获取用户id
	userID, err := controller.GetCurrentUser(c)
	if err != nil {
		controller.ResponseError(c, controller.CodeNeedLogin)
		return
	}
	if err := service.PostVote(userID, p); err != nil {
		zap.L().Error("service.PostVote(userID, p) failed", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, controller.CodeSuccess)
	return
}
