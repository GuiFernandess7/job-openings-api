package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}

func ReadOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}

func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}

func ListOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}

func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}

func ReadAllOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}