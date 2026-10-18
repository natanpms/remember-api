package usecase

import (
	"errors"

	"github.com/natanpms/remember-api/auth"
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repository repository.UserRepository
}

func NewAuthUseCase(repo repository.UserRepository) AuthUseCase {
	return AuthUseCase{
		repository: repo,
	}
}

func (au *AuthUseCase) Register(request model.RegisterRequest) (model.User, error) {
	_, err := au.repository.GetUserByEmail(request.Email)

	if err == nil {
		return model.User{}, errors.New("email já cadastrado")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	}

	return au.repository.CreateUser(user)
}

func (au *AuthUseCase) Login(request model.LoginRequest) (string, error) {
	user, err := au.repository.GetUserByEmail(request.Email)

	if err != nil {
		return "", errors.New("email ou senha inválidos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))

	if err != nil {
		return "", errors.New("email ou senha inválidos")
	}

	return auth.GenerateToken(user.ID)
}
