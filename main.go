package main

import (
	"blog_project/config"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetUpDatabaseConnection()
)

func main() {
	defer config.CloseDatabase(db)

}
