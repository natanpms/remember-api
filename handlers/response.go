package handlers

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code ErrorCode `json:"code"`
}

type ErrorCode string

const (
	ErrCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrCodeInvalidRequest ErrorCode = "INVALID_REQUEST"
	ErrCodeUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden      ErrorCode = "FORBIDDEN"
	ErrCodeInternal       ErrorCode = "INTERNAL_ERROR"
)

func Success(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, status int, code ErrorCode, message string) {
	ctx.JSON(status, APIResponse{
		Success: false,
		Message: message,
		Error: &APIError{
			Code: code,
		},
	})
}

func NotFound(ctx *gin.Context, message string) {
	Error(ctx, 404, ErrCodeNotFound, message)
}

func InvalidRequest(ctx *gin.Context, message string) {
	Error(ctx, 400, ErrCodeInvalidRequest, message)
}

func Unauthorized(ctx *gin.Context, message string) {
	Error(ctx, 401, ErrCodeUnauthorized, message)
}

func Forbidden(ctx *gin.Context, message string) {
	Error(ctx, 403, ErrCodeForbidden, message)
}

func InternalError(ctx *gin.Context, message string) {

	if message == "" {
		message = "Erro do servidor."
	}

	Error(ctx, 500, ErrCodeInternal, message)
}
