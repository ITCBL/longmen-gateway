package dto

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserLoginRequestParams struct {
	Mobile    string `json:"mobile" form:"mobile" comment:"电话号码" example:"13800138000" binding:"required,valid_mobile"` //todo 手机参数校验
	Password  string `json:"password" form:"password" comment:"密码" example:"123456" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type UserRegisterRequestParams struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,valid_mobile"` // 手机号码格式有规范可寻，自定义validater
	PassWord string `form:"password" json:"password" binding:"required,min=8,max=20"`
	//Code     string `form:"code" json:"code" binding:"required,min=6,max=6"` // todo 等开通短信服务再启用
}

type UserSession struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
