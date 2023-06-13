package repositories

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

type RefJafungRepository interface {
	FindAll() ([]migrations.RefJabFung, error)
	FindAllActive() ([]migrations.RefJabFung, error)
	FindAllPaginated(offset int, limit int) ([]migrations.RefJabFung, error)
	CountActive() int64
}

type repository struct {
	database *gorm.DB
}

func NewRefJafungRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAll() ([]migrations.RefJabFung, error) {
	var jafungs []migrations.RefJabFung
	err := r.database.Debug().Find(&jafungs).Error
	return jafungs, err
}

func (r *repository) FindAllActive() ([]migrations.RefJabFung, error) {
	var jafungs []migrations.RefJabFung
	err := r.database.Debug().Where("is_deleted = ?", 0).Find(&jafungs).Error
	return jafungs, err
}

func (r *repository) FindAllPaginated(offset int, limit int) ([]migrations.RefJabFung, error) {
	var jafungs []migrations.RefJabFung
	err := r.database.Debug().Where("is_deleted = ?", 0).Limit(limit).Offset(offset).Find(&jafungs).Error
	return jafungs, err
}

func (r *repository) CountActive() int64 {
	var count int64
	_ = r.database.Model(&migrations.RefJabFung{}).Where("is_deleted = ?", 0).Count(&count).Error
	return count
}
