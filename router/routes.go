package router

import (
	handler "github.com/GuiFernandess7/job-openings-api/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ReadOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}