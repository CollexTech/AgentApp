package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "message": "ok"})
}

func OptionsHandle(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}
