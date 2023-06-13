package services

import (
	"backend/models/migrations"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"
)

type RefJafungService interface {
	GetAll() ([]migrations.RefJabFung, error)
	GetAllActive() ([]returns.RefJabFung, error)
	GetAllPaginated(page int) (universals.Pagination, error)
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

func (s *service) GetAllPaginated(page int) (universals.Pagination, error) {
	var refJafungs []returns.RefJabFung
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10
	limit := 10
	per_page := 10
	count := s.repository.CountActive()
	dbRefJafungs, err := s.repository.FindAllPaginated(offset, limit)
	nomor := offset + 1
	for _, Q := range dbRefJafungs {
		refJafung := returns.RefJabFung{
			Kode:         Q.Kode,
			Nomor:        nomor,
			Nama_jab:     Q.Nama_jab,
			Nama_jenjang: Q.Nama_jenjang,
		}
		refJafungs = append(refJafungs, refJafung)
		nomor++
	}
	refJafungPagination := universals.Pagination{
		Current_page: page,
		Total:        count,
		Per_page:     per_page,
		Data:         refJafungs,
	}
	return refJafungPagination, err
}
