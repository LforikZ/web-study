package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExit
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeNeedLogin
	CodeComIDExit
	CodeComNameExit
	CodeGetListFiled
	CodeGetComFiled
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExit:        "用户名已存在",
	CodeUserNotExit:     "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效token",
	CodeNeedLogin:       "需要登录",
	CodeComIDExit:       "社区ID已存在",
	CodeComNameExit:     "社区名字已存在",
	CodeGetListFiled:    "获取数组失败",
	CodeGetComFiled:     "根据id获取社区失败",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
