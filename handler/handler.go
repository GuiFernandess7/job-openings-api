package handler

import (
	"github.com/GuiFernandess7/job-openings-api/config"
	"gorm.io/gorm"
)

var (
	logger  *config.Logger
	db		*gorm.DB
)

func InitializeHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}