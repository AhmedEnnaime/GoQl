package db

import (
	"log"

	"github.com/AhmedEnnaime/GoQl/config"
	logUtils "github.com/AhmedEnnaime/GoQl/utils/logger"
	stringUtils "github.com/AhmedEnnaime/GoQl/utils/stringify"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Model *gorm.DB

func ConnectDatabase() {
	var err error
	DbEnvVariables := config.DbEnvConfig

	dsn := stringUtils.CreateFormattedString("host=%s user=%s password=%s dbname=%s port=%s", DbEnvVariables.DbHost, DbEnvVariables.DbUser, DbEnvVariables.DbPassword, DbEnvVariables.DbName, DbEnvVariables.DbPort)
	Model, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logUtils.Error(err, "Error connecting to database: ")
	}

	log.Println("Connection Successful")
}
