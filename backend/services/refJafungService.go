package services

import (
	"backend/models/migrations"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"
)

func NewRefJafungService(repository repositories.RefJafungRepository) *service {
	return &service{repository}
}

func (s *service) GetAllJafung() ([]migrations.RefJabFung, error) {
	return s.repository.FindAllJafung()
}

func (s *service) GetAllActiveJafung() ([]returns.RefJabFung, error) {
	var refJafungs []returns.RefJabFung
	dbRefJafungs, err := s.repository.FindAllActiveJafung()
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

func (s *service) GetAllPaginatedJafung(page int) (universals.Pagination, error) {
	var refJafungs []returns.RefJabFung
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10
	limit := 10
	per_page := 10
	count := s.repository.CountActiveJafung()
	dbRefJafungs, err := s.repository.FindAllPaginatedJafung(offset, limit)
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
