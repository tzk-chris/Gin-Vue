package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tzk-chris/Gin-Vue/controller"
	"github.com/tzk-chris/Gin-Vue/middleware"
)

func CollectRoute(router *gin.Engine) *gin.Engine {
	router.POST("/api/auth/register", controller.Register)
	router.POST("/api/auth/login", controller.Login)
	router.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return router
}
