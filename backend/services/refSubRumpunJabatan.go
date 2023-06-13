package services

import (
	"backend/models/migrations"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"
)

func NewRefSubRumpunJabatanService(refSubRumpunJabatanRepository repositories.RefSubRumpunJabatanRepository) *refSubRumpunJabatanService {
	return &refSubRumpunJabatanService{refSubRumpunJabatanRepository}
}

func (s *refSubRumpunJabatanService) GetAllSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error) {
	return s.refSubRumpunJabatanRepository.FindAllRefSubRumpunJabatan()
}

func (s *refSubRumpunJabatanService) GetAllActiveSubRumpunJabatan() ([]returns.RefSubrumpunJabatan, error) {
	var refSubRumpunJabatans []returns.RefSubrumpunJabatan
	dbRefSubRumpunJabatans, err := s.refSubRumpunJabatanRepository.FindAllActiveRefSubRumpunJabatan()
	for _, Q := range dbRefSubRumpunJabatans {
		refSubRumpunJabatan := returns.RefSubrumpunJabatan{
			Id_subrumpun:     Q.Id_subrumpun,
			Id_rumpun:        Q.Id_rumpun,
			Subrumpun:        Q.Subrumpun,
			Definisi:         Q.Definisi,
			Perangkat_daerah: Q.Perangkat_daerah,
		}
		refSubRumpunJabatans = append(refSubRumpunJabatans, refSubRumpunJabatan)
	}
	return refSubRumpunJabatans, err
}

func (s *refSubRumpunJabatanService) GetAllPaginatedSubRumpunJabatan(page int) (universals.Pagination, error) {
	var refSubRumpunJabatans []returns.RefSubrumpunJabatan
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10
	limit := 10
	per_page := 10
	count := s.refSubRumpunJabatanRepository.CountActiveRefSubRumpunJabatan()
	dbRefSubRumpunJabatans, err := s.refSubRumpunJabatanRepository.FindAllPaginatedRefSubRumpunJabatan(offset, limit)
	nomor := offset + 1
	for _, Q := range dbRefSubRumpunJabatans {
		refSubRumpunJabatan := returns.RefSubrumpunJabatan{
			Id_subrumpun:     Q.Id_subrumpun,
			Nomor:            nomor,
			Id_rumpun:        Q.Id_rumpun,
			Subrumpun:        Q.Subrumpun,
			Definisi:         Q.Definisi,
			Perangkat_daerah: Q.Perangkat_daerah,
		}
		refSubRumpunJabatans = append(refSubRumpunJabatans, refSubRumpunJabatan)
		nomor++
	}
	refSubRumpunJabatanPagination := universals.Pagination{
		Current_page: page,
		Total:        count,
		Per_page:     per_page,
		Data:         refSubRumpunJabatans,
	}
	return refSubRumpunJabatanPagination, err
}
