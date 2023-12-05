package routes

import (
	"github.com/Abdul-Rahim-K-T/pipeline/handler"
	"github.com/gin-gonic/gin"
)

func Router(ctx *gin.Engine){
	ctx.POST("create-user",handler.CreateUser)
	ctx.GET("get-user/:id",handler.GetUserById)
	ctx.PUT("update-user/:id",handler.UpdateUserById)
}