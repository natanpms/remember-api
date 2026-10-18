package routes

import (
	"database/sql"

	"github.com/natanpms/remember-api/controller"
	_ "github.com/natanpms/remember-api/docs"
	response "github.com/natanpms/remember-api/handlers"
	"github.com/natanpms/remember-api/repository"
	"github.com/natanpms/remember-api/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ExecuteRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) {
		response.NotFound(ctx, "Rota não existente na API.")
	})

	ProductRepository := repository.NewProductRepository(db)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	api := router.Group("/api/v1")

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.GET("/products", ProductController.GetProducts)
	api.GET("/products/:id", ProductController.GetProductByID)
	api.POST("/products", ProductController.CreateProduct)
	api.DELETE("/products/:id", ProductController.DeleteProduct)

	return router
}
