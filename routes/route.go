package routes

import (
	_ "gin-quickstart/docs"
	"net/http"

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
		"message": "pong teste",
	})
}

func ExecuteRoutes() *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "ROUTE_NOT_FOUND",
			"message": "A rota solicitada não existe.",
		})
	})

	api := router.Group("/api/v1")

	api.GET("/ping", Ping)
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
