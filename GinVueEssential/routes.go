package main

import (
	"github.com/gin-gonic/gin"
	"herrkung.com/GinVueEssential/controller"
	"herrkung.com/GinVueEssential/middleware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update) // 跟patch的区别就是，update是替换，patch是部分修改
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	//categoryRoutes.PATCH("/:id", )  // 部分修改

	return r
}
