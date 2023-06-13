package repositories

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

func NewRefSubRumpunJabatanRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAllRefSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error) {
	var subRumpunJabatans []migrations.RefSubrumpunJabatan
	err := r.database.Debug().Find(&subRumpunJabatans).Error
	return subRumpunJabatans, err
}

func (r *repository) FindAllActiveRefSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error) {
	var subRumpunJabatans []migrations.RefSubrumpunJabatan
	err := r.database.Debug().Where("is_deleted = ?", 0).Find(&subRumpunJabatans).Error
	return subRumpunJabatans, err
}

func (r *repository) FindAllPaginatedRefSubRumpunJabatan(offset int, limit int) ([]migrations.RefSubrumpunJabatan, error) {
	var subRumpunJabatans []migrations.RefSubrumpunJabatan
	err := r.database.Debug().Where("is_deleted = ?", 0).Limit(limit).Offset(offset).Find(&subRumpunJabatans).Error
	return subRumpunJabatans, err
}

func (r *repository) CountActiveRefSubRumpunJabatan() int64 {
	var count int64
	_ = r.database.Model(&migrations.RefSubrumpunJabatan{}).Where("is_deleted = ?", 0).Count(&count).Error
	return count
}
