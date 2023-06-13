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
	GetAllRumpunJabatan() ([]migrations.RefJabFung, error)
	GetAllActiveRumpunJabatan() ([]returns.RefJabFung, error)
	GetAllPaginatedRumpunJabatan(page int) (universals.Pagination, error)
}

type refSubRumpunJabatanService struct {
	refSubRumpunJabatanRepository repositories.RefSubRumpunJabatanRepository
}

type RefSubRumpunJabatanService interface {
	GetAllSubRumpunJabatan() ([]migrations.RefJabFung, error)
	GetAllActiveSubRumpunJabatan() ([]returns.RefJabFung, error)
	GetAllPaginatedSubRumpunJabatan(page int) (universals.Pagination, error)
}
