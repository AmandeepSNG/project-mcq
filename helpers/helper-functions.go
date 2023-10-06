package helpers

import (
	"Plateform-MCQ/common"

	"github.com/gin-gonic/gin"
)

func GetCurrentEnv(ENV string) string {
	switch ENV {
	case "dev":
		return common.ENV_LOCAL
	case "local":
		return common.ENV_LOCAL
	case "staging":
		return common.ENV_STAGING
	case "prod":
		return common.ENV_PRODUCTION
	case "production":
		return common.ENV_PRODUCTION
	default:
		return common.ENV_LOCAL
	}
}

func GenerateErrorResponse(ctx *gin.Context, message string, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
		"error":   err.Error(),
	})
}

func GenerateSuccessResponse(ctx *gin.Context, message string, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
		"data":    data,
	})
}
