package auth

import (
	"strings"

	"github.com/natanpms/remember-api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			handlers.Unauthorized(ctx, "Token não informado")
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		userID, err := ValidateToken(tokenString)

		if err != nil {
			handlers.Unauthorized(ctx, "Token inválido ou expirado")
			ctx.Abort()
			return
		}

		ctx.Set("userID", userID)
		ctx.Next()
	}
}
