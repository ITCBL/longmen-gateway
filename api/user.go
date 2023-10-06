package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"longmen-gateway/dto"
	"longmen-gateway/global"
	"longmen-gateway/middlewares"
	"longmen-gateway/model"
	"net/http"
	"time"
)

type UserApi struct{}

func (u *UserApi) Login(ctx *gin.Context) {
	params := &dto.UserLoginRequestParams{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		middlewares.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// 用户是否存在
	var user model.User
	userInfo, _ := user.GetUserByMobile(params.Mobile)
	if userInfo == nil {
		middlewares.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("用户不存在"))
		return
	}

	// 密码校验
	isPass, _ := user.CheckPassword(params.Password, userInfo.Password)

	if isPass { // 密码匹配通过
		// 生成token
		newJWT := middlewares.NewJWT()
		customClaims := dto.CustomClaims{
			ID:          uint(userInfo.ID),
			NickName:    userInfo.NickName,
			AuthorityId: uint(userInfo.Role),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.JwtTokenExpirationTime, // 30天过期
				Issuer:    global.Issuer,
				NotBefore: time.Now().Unix(),
			},
		}
		token, err := newJWT.CreateToken(customClaims)
		if err != nil {
			middlewares.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("生成token失败"))
			return
		}
		middlewares.SuccessResponse(ctx, gin.H{"id": userInfo.ID,
			"nick_name":  userInfo.NickName,
			"token":      token,
			"expires_at": (time.Now().Unix() + global.JwtTokenExpirationTime) * 1000,
		})
		return
	}

	middlewares.ErrorResponse(ctx, http.StatusBadRequest, errors.New("密码错误"))
}

func (u *UserApi) Register(ctx *gin.Context) {
	requestParams := dto.UserRegisterRequestParams{}
	if err := ctx.ShouldBindJSON(&requestParams); err != nil {
		middlewares.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// 验证码 todo 需要到短信平台开通，本功能暂时略过

	var user model.User

	// 用户是否存在
	userInfo, err := user.GetUserByMobile(requestParams.Mobile)
	if userInfo != nil {
		middlewares.ErrorResponse(ctx, http.StatusBadRequest, errors.New("用户已存在"))
		return
	}

	// 创建用户
	createUser, err := user.CreateUser(requestParams.Mobile, requestParams.PassWord, requestParams.Mobile)
	if err != nil {
		middlewares.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("创建用户失败"))
		return
	}

	// 注册成功并登陆
	newJWT := middlewares.NewJWT()
	customClaims := dto.CustomClaims{
		ID:          uint(createUser.ID),
		NickName:    createUser.NickName,
		AuthorityId: uint(createUser.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(), //签名的生效时间
			Issuer:    global.Issuer,
			NotBefore: time.Now().Unix() + global.JwtTokenExpirationTime, //30天过期
		},
	}
	token, err := newJWT.CreateToken(customClaims)
	if err != nil {
		middlewares.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("生成token失败"))
		return
	}

	middlewares.SuccessResponse(ctx, gin.H{
		"id":         createUser.ID,
		"nick_name":  createUser.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func (u *UserApi) Logout(ctx *gin.Context) { // todo

}

func (u *UserApi) GetUserInfo(ctx *gin.Context) { // todo

}

func (u *UserApi) UpdatePwd(ctx *gin.Context) { // todo

}
