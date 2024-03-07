package main

import (
	"github.com/gin-gonic/gin"
	"jwt-gin-example/controllers"
	"jwt-gin-example/middlewares"
	"jwt-gin-example/models"
)

func main() {

	defer func() {
		models.Cleanup()
	}()

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	// 以下接口将需要使用jwt进行认证
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	_ = r.Run(":8080")

}
