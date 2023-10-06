package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"longmen-gateway/middlewares"
	"net/http"
)

// 存储
var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	// 生成驱动
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	newCaptcha := base64Captcha.NewCaptcha(driver, store) // 驱动+存储

	// 生成验证码
	id, b64s, err := newCaptcha.Generate()
	if err != nil {
		middlewares.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("生成验证码错误"))
		return
	}

	middlewares.SuccessResponse(ctx, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
