package routes

import (
	_ "github.com/natanpms/remember-api/docs"
	response "github.com/natanpms/remember-api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Ping godoc
// @Summary ping example
// @Description do ping
// @Tags ping
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ExecuteRoutes() *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) {
		response.NotFound(ctx, "Rota não existente na API.")
	})

	api := router.Group("/api/v1")

	authRoutes := api.Group("/auth")
	authRoutes.POST("/login")
	authRoutes.POST("/register")

	api.GET("/ping", Ping)
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
