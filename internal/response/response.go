package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, message string, details any) {
	response := gin.H{
		"message": message,
		"code":    code,
	}

	if details != nil {
		response["details"] = details
	}

	ctx.JSON(code, response)
}

func SendSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
