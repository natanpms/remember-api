package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/natanpms/remember-api/handlers"
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/usecase"
)

type authController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(usecase usecase.AuthUseCase) authController {
	return authController{
		authUseCase: usecase,
	}
}

// Register
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "User payload"
// @Success 201 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Router /api/v1/auth/register [post]
func (ac *authController) Register(ctx *gin.Context) {
	var request model.RegisterRequest

	err := ctx.BindJSON(&request)

	if err != nil {
		handlers.InvalidRequest(ctx, "Erro ao ler dados do usuário")
		return
	}

	user, err := ac.authUseCase.Register(request)

	if err != nil {
		handlers.Error(ctx, 400, handlers.ErrCodeInvalidRequest, err.Error())
		return
	}

	handlers.Success(ctx, 201, "Usuário criado com sucesso", user)
}

// Login
// @Summary Login
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "Login credentials"
// @Success 200 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Router /api/v1/auth/login [post]
func (ac *authController) Login(ctx *gin.Context) {
	var request model.LoginRequest

	err := ctx.BindJSON(&request)

	if err != nil {
		handlers.InvalidRequest(ctx, "Erro ao ler dados de login")
		return
	}

	token, err := ac.authUseCase.Login(request)

	if err != nil {
		handlers.Unauthorized(ctx, err.Error())
		return
	}

	handlers.Success(ctx, 200, "Login realizado com sucesso", gin.H{"token": token})
}
