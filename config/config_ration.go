package config

import (
	"os"

	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	dbUser := os.Getenv("DB_USER")

}
