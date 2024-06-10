package router

import "github.com/gin-gonic/gin"

func Initialize(){
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	router.Run(":8080")
}

// https://www.youtube.com/watch?v=wyEYpX5U4Vg