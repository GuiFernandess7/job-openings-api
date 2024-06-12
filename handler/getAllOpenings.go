package handler

import (
	"net/http"
	"strconv"

	"github.com/GuiFernandess7/job-openings-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListOpeningHandler(ctx *gin.Context) {
    var limit int
    openings := []schemas.Opening{}
    limitStr := ctx.Query("limit")

    if limitStr != "" {
        parsedLimit, err := strconv.Atoi(limitStr)
        if err != nil {
            sendError(ctx, http.StatusBadRequest, "invalid limit parameter")
            return
        }
        limit = parsedLimit
        query := db.Model(&schemas.Opening{}).Order("created_at DESC").Limit(limit)

        if err := sendData(query, &openings); err != nil {
            sendError(ctx, http.StatusInternalServerError, "error listing openings")
            return
        }
    } else {
        if err := sendData(db, &openings); err != nil {
            sendError(ctx, http.StatusInternalServerError, "error listing openings")
            return
        }
    }

    sendSuccess(ctx, "list-openings", openings)
}

func sendData(engine *gorm.DB, destinations interface{}) error {
    if err := engine.Find(destinations).Error; err != nil {
        return err
    }
    return nil
}
