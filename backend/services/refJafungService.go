package services

import (
	"backend/models/migrations"
	"backend/models/returns"
	"backend/repositories"
	"backend/universalfunctions"
)

type RefJafungService interface {
	GetAll() ([]migrations.RefJabFung, error)
	GetAllActive() ([]returns.RefJabFung, error)
	GetAllPaginated(pagination universalfunctions.Pagination) (*universalfunctions.Pagination, error)
}

type service struct {
	repository repositories.RefJafungRepository
}

func NewRefJafungService(repository repositories.RefJafungRepository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]migrations.RefJabFung, error) {
	return s.repository.FindAll()
}

func (s *service) GetAllActive() ([]returns.RefJabFung, error) {
	var refJafungs []returns.RefJabFung
	dbRefJafungs, err := s.repository.FindAllActive()
	for _, Q := range dbRefJafungs {
		refJafung := returns.RefJabFung{
			Kode:         Q.Kode,
			Nama_jab:     Q.Nama_jab,
			Nama_jenjang: Q.Nama_jenjang,
		}
		refJafungs = append(refJafungs, refJafung)
	}
	return refJafungs, err
}

func (s *service) GetAllPaginated(pagination universalfunctions.Pagination) (*universalfunctions.Pagination, error) {
	return s.repository.FindAllPaginated(pagination)
}
