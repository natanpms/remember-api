package routes

import (
	"database/sql"

	"github.com/natanpms/remember-api/auth"
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

	UserRepository := repository.NewUserRepository(db)
	AuthUseCase := usecase.NewAuthUseCase(UserRepository)
	AuthController := controller.NewAuthController(AuthUseCase)

	MemoryGroupRepository := repository.NewMemoryGroupRepository(db)
	MemoryGroupUseCase := usecase.NewMemoryGroupUseCase(MemoryGroupRepository)
	MemoryGroupController := controller.NewMemoryGroupController(MemoryGroupUseCase)

	MemoryRepository := repository.NewMemoryRepository(db)
	MemoryUseCase := usecase.NewMemoryUseCase(MemoryRepository, MemoryGroupRepository)
	MemoryController := controller.NewMemoryController(MemoryUseCase)

	api := router.Group("/api/v1")

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes := api.Group("/auth")
	authRoutes.POST("/register", AuthController.Register)
	authRoutes.POST("/login", AuthController.Login)

	protected := api.Group("")
	protected.Use(auth.AuthMiddleware())

	protected.GET("/groups", MemoryGroupController.GetGroups)
	protected.GET("/groups/:id", MemoryGroupController.GetGroupByID)
	protected.POST("/groups", MemoryGroupController.CreateGroup)
	protected.DELETE("/groups/:id", MemoryGroupController.DeleteGroup)

	protected.GET("/groups/:id/memories", MemoryController.GetMemoriesByGroup)
	protected.POST("/groups/:id/memories", MemoryController.CreateMemory)
	protected.GET("/memories/:id", MemoryController.GetMemoryByID)
	protected.DELETE("/memories/:id", MemoryController.DeleteMemory)

	return router
}
