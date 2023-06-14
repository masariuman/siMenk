package repositories

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

func NewRefRumpunJabatanRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAllRefRumpunJabatan() ([]migrations.RefRumpunJabatan, error) {
	var rumpunJabatans []migrations.RefRumpunJabatan
	err := r.database.Debug().Find(&rumpunJabatans).Error
	return rumpunJabatans, err
}

func (r *repository) FindAllActiveRefRumpunJabatan() ([]migrations.RefRumpunJabatan, error) {
	var rumpunJabatans []migrations.RefRumpunJabatan
	err := r.database.Debug().Find(&rumpunJabatans).Error
	return rumpunJabatans, err
}

func (r *repository) FindAllPaginatedRefRumpunJabatan(offset int, limit int) ([]migrations.RefRumpunJabatan, error) {
	var rumpunJabatans []migrations.RefRumpunJabatan
	err := r.database.Debug().Limit(limit).Offset(offset).Find(&rumpunJabatans).Error
	return rumpunJabatans, err
}

func (r *repository) CountActiveRefRumpunJabatan() int64 {
	var count int64
	_ = r.database.Model(&migrations.RefRumpunJabatan{}).Count(&count).Error
	return count
}
