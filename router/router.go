package router

import "github.com/gin-gonic/gin"

func Initialize(){
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8080")
}

// https://www.youtube.com/watch?v=wyEYpX5U4Vg