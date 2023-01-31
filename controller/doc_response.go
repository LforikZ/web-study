// @Author Zihao_Li 2023/1/31 15:51:00
package controller

type _ResponseVote struct {
	Code    ResCode `json:"code"`    //业务响应码
	Message string  `json:"message"` //提示信息
	Data    string  `json:"data"`    //数据
}
