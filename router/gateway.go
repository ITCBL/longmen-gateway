package router

import "github.com/gin-gonic/gin"

func GatewayRouter(group *gin.RouterGroup) {
	group.Group("/gateway", nil)
}
