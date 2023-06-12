package configs

import (
	"gorm.io/gorm"
)

func autoMigration(connect *gorm.DB) {
	// connect.AutoMigrate(migrations.RefJabFung{})
}
