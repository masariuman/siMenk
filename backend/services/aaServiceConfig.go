package services

import (
	"backend/models/migrations"
	"backend/models/requests"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"

	"github.com/gin-gonic/gin"
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
	Store(refRumpunJabatanRequest requests.RefRumpunJabatanRequest, err error, c *gin.Context) (returns.RefRumpunJabatan, error)
	GetById(ID int) (returns.RefRumpunJabatan, error)
	Update(ID int, refRumpunJabatanRequest requests.RefRumpunJabatanRequest) (error)
}

type refSubRumpunJabatanService struct {
	refSubRumpunJabatanRepository repositories.RefSubRumpunJabatanRepository
}

type RefSubRumpunJabatanService interface {
	GetAllSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error)
	GetAllActiveSubRumpunJabatan() ([]returns.RefSubrumpunJabatan, error)
	GetAllPaginatedSubRumpunJabatan(page int) (universals.Pagination, error)
}
