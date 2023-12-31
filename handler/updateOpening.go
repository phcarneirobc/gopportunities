package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phcarneirobc/gopportunities/schemas"
)	
func UpdateOpeningHandler(ctx *gin.Context){
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil{
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id  := ctx.Query("id")
	if id == ""{
		sendError(ctx,http.StatusBadRequest, errParamIsRequired("id","queryParameter").Error())
		return
	}
	opening := schemas.Opening{}


	if err := db.First(&opening , id).Error; err != nil{
		sendError(ctx,http.StatusNotFound,"opening not found")
		return
	}
	if request.Name != "" {
		opening.Name = request.Name
	}
	if request.CPF != "" {
		opening.CPF = request.CPF
	}
	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Priority != nil {
		opening.Priority = *request.Priority
	}

	if request.Number != "" {
		opening.Number = request.Number
	}
	/*
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	*/
	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSucess(ctx, "update-opening", opening)
}