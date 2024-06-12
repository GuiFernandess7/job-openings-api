package handler

import (
	"net/http"

	"github.com/GuiFernandess7/job-openings-api/schemas"
	"github.com/gin-gonic/gin"
)

func ReadOpeningHandler(ctx *gin.Context){
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
	} else {
		sendSuccess(ctx, "show-opening", opening)
	}
}

