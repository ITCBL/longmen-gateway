package router

import "github.com/gin-gonic/gin"

func ManagerRouter(group *gin.RouterGroup) {
	group.Group("/manager", nil)
}
