package repositories

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

type RefJafungRepository interface {
	FindAll() ([]migrations.RefJabFung, error)
}

type repository struct {
	database *gorm.DB
}

func NewRefJafungRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAll() ([]migrations.RefJabFung, error) {
	var artikels []migrations.RefJabFung
	err := r.database.Debug().Find(&artikels).Error
	return artikels, err
}
