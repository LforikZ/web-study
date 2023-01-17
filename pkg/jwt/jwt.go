package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration token过期时间 2小时
const TokenExpireDuration = time.Hour * 2

// MySecret 盐值
var MySecret = []byte("李子爱猪猪")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(username string, userID int64) (aToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID, // 自定义字段
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间  2小时
			Issuer:    "web-study",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象 atoken
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)

	//TODO: 需要rToken时放开
	//// refresh token 不需要任何自定义数据
	//rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Second * 30).Unix(), //过期时间 30天
	//	Issuer:    "web-study",                             //签发人
	//}).SignedString(MySecret)

	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, Keyfunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid { //校验token
		err = errors.New("invalid token")
	}
	return
}

// TODO: 这个功能没有实现，学习完之后需要自己完成
func RefreshToken(aToken, rToken string) (newAToken string, err error) {
	// refresh token 无效直接返回
	if _, err = jwt.Parse(rToken, Keyfunc); err != nil {
		return
	}

	//从旧的access token 中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, Keyfunc)
	v, _ := err.(*jwt.ValidationError)

	//当access token是过期错误 并且 refresh token没有过期时 就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.Username, claims.UserId)
	}
	return
}
func Keyfunc(token *jwt.Token) (interface{}, error) {
	return "", errors.New("123")
}
