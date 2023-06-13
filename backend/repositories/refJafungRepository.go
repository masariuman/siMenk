package repositories

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

func NewRefJafungRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAllJafung() ([]migrations.RefJabFung, error) {
	var jafungs []migrations.RefJabFung
	err := r.database.Debug().Find(&jafungs).Error
	return jafungs, err
}

func (r *repository) FindAllActiveJafung() ([]migrations.RefJabFung, error) {
	var jafungs []migrations.RefJabFung
	err := r.database.Debug().Where("is_deleted = ?", 0).Find(&jafungs).Error
	return jafungs, err
}

func (r *repository) FindAllPaginatedJafung(offset int, limit int) ([]migrations.RefJabFung, error) {
	var jafungs []migrations.RefJabFung
	err := r.database.Debug().Where("is_deleted = ?", 0).Limit(limit).Offset(offset).Find(&jafungs).Error
	return jafungs, err
}

func (r *repository) CountActiveJafung() int64 {
	var count int64
	_ = r.database.Model(&migrations.RefJabFung{}).Where("is_deleted = ?", 0).Count(&count).Error
	return count
}
