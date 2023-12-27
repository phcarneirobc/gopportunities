package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phcarneirobc/gopportunities/schemas"
)
func CreateOpeningHandler(ctx *gin.Context){
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	
	if err :=	request.Validate() ; err != nil{
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	
	opening := schemas.Opening{
		Name: request.Name,
		CPF: request.CPF,
		Location: request.Location,
		Priority: *request.Priority,
		Number: request.Number,
		//Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil{
		logger.Errorf("error creating opening: %v",err.Error())
		sendError(ctx,http.StatusInternalServerError, "error creating opening on database")
		return
	}


	sendSucess(ctx,"create-opening",opening)
}

