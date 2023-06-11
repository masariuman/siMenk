package configs

import (
	"backend/universalfunctions"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connect *gorm.DB

func ConnectDatabase() {
	USERNAME := "root"
	PASSWORD := ""
	HOST := "127.0.0.1"
	PORT := "3306"
	DATABASE := "simpegme_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DATABASE)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	universalfunctions.PanicErr(err)

	// autoMigration(database)
	Connect = database
}
