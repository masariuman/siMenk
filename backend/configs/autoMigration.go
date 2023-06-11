package configs

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

func autoMigration(connect *gorm.DB) {
	connect.AutoMigrate(migrations.RefJabFung{})
}
