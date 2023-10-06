package router

import (
	"github.com/gin-gonic/gin"
	"longmen-gateway/api"
)

func ManagerRouter(group *gin.RouterGroup) {
	managerGroup := group.Group("/manager")

	// 用户模块
	{
		userGroup := managerGroup.Group("/user")
		userApi := &api.UserApi{}
		userGroup.POST("/login", userApi.Login)         // 登录
		userGroup.POST("/register", userApi.Register)   // 注册
		userGroup.GET("/logout", userApi.Logout)        // 退出 todo
		userGroup.GET("/info/:id", userApi.GetUserInfo) // 获取用户详情 todo
		userGroup.PUT("/pwd/:id", userApi.UpdatePwd)    // 修改密码 todo
	}

	{
		baseGroup := managerGroup.Group("/base")
		baseGroup.GET("captcha", api.GetCaptcha) // 图片验证码
		baseGroup.POST("send_sms", api.SendSms)  // 短信验证码 todo
	}

	// 仪表盘
	{

	}

	// 服务管理
	{

	}
}
