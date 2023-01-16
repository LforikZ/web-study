package entity

// ParamSignUp 定义请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type ParamLoginUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
