package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbConfig := AppConfigData.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.Charset,
		dbConfig.Loc,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB = database
	fmt.Println("Database connection established")
}
