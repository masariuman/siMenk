package services

import (
	"backend/models/migrations"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"
)

func NewRefRumpunJabatanService(refRumpunJabatanRepository repositories.RefRumpunJabatanRepository) *refRumpunJabatanservice {
	return &refRumpunJabatanservice{refRumpunJabatanRepository}
}

func (s *refRumpunJabatanservice) GetAllRumpunJabatan() ([]migrations.RefRumpunJabatan, error) {
	return s.refRumpunJabatanRepository.FindAllRefRumpunJabatan()
}

func (s *refRumpunJabatanservice) GetAllActiveRumpunJabatan() ([]returns.RefRumpunJabatan, error) {
	var refRumpunJabatans []returns.RefRumpunJabatan
	dbRefRumpunJabatans, err := s.refRumpunJabatanRepository.FindAllActiveRefRumpunJabatan()
	for _, Q := range dbRefRumpunJabatans {
		refRumpunJabatan := returns.RefRumpunJabatan{
			Id_rumpun: Q.Id_rumpun,
			Rumpun:    Q.Rumpun,
		}
		refRumpunJabatans = append(refRumpunJabatans, refRumpunJabatan)
	}
	return refRumpunJabatans, err
}

func (s *refRumpunJabatanservice) GetAllPaginatedRumpunJabatan(page int) (universals.Pagination, error) {
	var refRumpunJabatans []returns.RefRumpunJabatan
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10
	limit := 10
	per_page := 10
	count := s.refRumpunJabatanRepository.CountActiveRefRumpunJabatan()
	dbRefRumpunJabatans, err := s.refRumpunJabatanRepository.FindAllPaginatedRefRumpunJabatan(offset, limit)
	nomor := offset + 1
	for _, Q := range dbRefRumpunJabatans {
		refRumpunJabatan := returns.RefRumpunJabatan{
			Id_rumpun: Q.Id_rumpun,
			Nomor:     nomor,
			Rumpun:    Q.Rumpun,
		}
		refRumpunJabatans = append(refRumpunJabatans, refRumpunJabatan)
		nomor++
	}
	refRumpunJabatanPagination := universals.Pagination{
		Current_page: page,
		Total:        count,
		Per_page:     per_page,
		Data:         refRumpunJabatans,
	}
	return refRumpunJabatanPagination, err
}
