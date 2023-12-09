package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "GET OPENING",
		})
}

func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"msg": "GET OPENING",
	})
}

func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"msg": "GET OPENING",
	})
}
func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"msg": "GET OPENING",
	})
}
func ListOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"msg": "GET OPENING",
	})
}