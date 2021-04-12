package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tzk-chris/Gin-Vue/controller"
)

func CollectRoute(router *gin.Engine) *gin.Engine {
	router.POST("/api/auth/register", controller.Register)

	return router
}
