package services

import (
	"backend/models/migrations"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"
)

type service struct {
	repository repositories.RefJafungRepository
}

type RefJafungService interface {
	GetAllJafung() ([]migrations.RefJabFung, error)
	GetAllActiveJafung() ([]returns.RefJabFung, error)
	GetAllPaginatedJafung(page int) (universals.Pagination, error)
}

type refRumpunJabatanservice struct {
	refRumpunJabatanRepository repositories.RefRumpunJabatanRepository
}

type RefRumpunJabatanService interface {
	GetAllRumpunJabatan() ([]migrations.RefRumpunJabatan, error)
	GetAllActiveRumpunJabatan() ([]returns.RefRumpunJabatan, error)
	GetAllPaginatedRumpunJabatan(page int) (universals.Pagination, error)
}

type refSubRumpunJabatanService struct {
	refSubRumpunJabatanRepository repositories.RefSubRumpunJabatanRepository
}

type RefSubRumpunJabatanService interface {
	GetAllSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error)
	GetAllActiveSubRumpunJabatan() ([]returns.RefSubrumpunJabatan, error)
	GetAllPaginatedSubRumpunJabatan(page int) (universals.Pagination, error)
}
