package services

import (
	"backend/models/migrations"
	"backend/models/requests"
	"backend/models/returns"
	"backend/repositories"
	"backend/universals"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		nomorString := strconv.Itoa(nomor)
		refRumpunJabatan := returns.RefRumpunJabatan{
			Id_rumpun: Q.Id_rumpun,
			Nomor:     nomorString,
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

func (s *refRumpunJabatanservice) Store(refRumpunJabatanRequest requests.RefRumpunJabatanRequest, err error, c *gin.Context) (returns.RefRumpunJabatan, error) {
	if err != nil {
		var returnRefRumpunJabatan returns.RefRumpunJabatan
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error pada bagian %s, Kondisi : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return returnRefRumpunJabatan, err
	}
	refRumpunJabatan := migrations.RefRumpunJabatan{
		Rumpun: refRumpunJabatanRequest.Rumpun,
	}
	_, err = s.refRumpunJabatanRepository.Store(refRumpunJabatan)
	universals.PanicErr(err)
	refRumpunJabatanLast, _ := s.refRumpunJabatanRepository.FindLast()
	refRumpunJabatanLatest := returns.RefRumpunJabatan{
		Id_rumpun: refRumpunJabatanLast.Id_rumpun,
		Nomor:     "NEW",
		Rumpun:    refRumpunJabatanLast.Rumpun,
	}
	return refRumpunJabatanLatest, err
}

func (s *refRumpunJabatanservice) GetById(ID int) (returns.RefRumpunJabatan, error) {
	refRumpunJabatanFind, err := s.refRumpunJabatanRepository.FindById(ID)
	refRumpunJabatanFound := returns.RefRumpunJabatan{
		Id_rumpun: refRumpunJabatanFind.Id_rumpun,
		Nomor:     "EDIT",
		Rumpun:    refRumpunJabatanFind.Rumpun,
	}
	return refRumpunJabatanFound, err
}

func (s *refRumpunJabatanservice) Update(ID int, refRumpunJabatanRequest requests.RefRumpunJabatanRequest) error {
	refRumpunJabatan := migrations.RefRumpunJabatan{
		Id_rumpun: ID,
		Rumpun:    refRumpunJabatanRequest.Rumpun,
	}
	_, err := s.refRumpunJabatanRepository.Update(ID, refRumpunJabatan)
	return err
}
