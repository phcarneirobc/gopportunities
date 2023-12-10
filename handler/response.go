package handler

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int , msg string){
	ctx.Header("Context-type", "aplication/json")
	ctx.JSON(code, gin.H{
		"msg": msg,
		"errorCode": code,
	})
}

func sendSucess (ctx *gin.Context, op string, data interface{}){
	ctx.Header("Context-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message":fmt.Sprintf("operation from handler %s sucesseful", op) ,
		"data": data ,
	})
}