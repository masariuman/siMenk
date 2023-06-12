package repositories

import (
	"backend/models/migrations"
	"backend/universalfunctions"

	"gorm.io/gorm"
)

type RefJafungRepository interface {
	FindAll() ([]migrations.RefJabFung, error)
	FindAllActive() ([]migrations.RefJabFung, error)
	FindAllPaginated(pagination universalfunctions.Pagination) (*universalfunctions.Pagination, error)
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

func (r *repository) FindAllPaginated(pagination universalfunctions.Pagination) (*universalfunctions.Pagination, error) {
	var jafungs []migrations.RefJabFung
	r.database.Scopes(universalfunctions.Paginate(jafungs, &pagination, r.database)).Find(&jafungs)
	pagination.Rows = jafungs

	return &pagination, nil
}
