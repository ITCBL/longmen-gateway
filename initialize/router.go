package initialize

import (
	"github.com/gin-gonic/gin"
	"longmen-gateway/middlewares"
	"longmen-gateway/router"
	"net/http"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())
	Router.Use()
	ApiGroup := Router.Group("/api/v1")
	Router.NoRoute(NoRouteHandles)
	router.ManagerRouter(ApiGroup)
	return Router
}

func NoRouteHandles(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"msg": "The URL accessed does not exist!",
	})
}
