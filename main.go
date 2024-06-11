package main

import (
	"github.com/GuiFernandess7/job-openings-api/config"
	"github.com/GuiFernandess7/job-openings-api/router"
)

var (
	logger *config.Logger
)

func main(){
	logger = config.GetLogger("main")

	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	// Initialize router
	router.Initialize()
}

// https://www.youtube.com/watch?v=wyEYpX5U4Vg (1:52:11)