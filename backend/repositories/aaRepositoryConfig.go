package repositories

import (
	"backend/models/migrations"

	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
}

type RefJafungRepository interface {
	FindAllJafung() ([]migrations.RefJabFung, error)
	FindAllActiveJafung() ([]migrations.RefJabFung, error)
	FindAllPaginatedJafung(offset int, limit int) ([]migrations.RefJabFung, error)
	CountActiveJafung() int64
}

type RefRumpunJabatanRepository interface {
	FindAllRefRumpunJabatan() ([]migrations.RefRumpunJabatan, error)
	FindAllActiveRefRumpunJabatan() ([]migrations.RefRumpunJabatan, error)
	FindAllPaginatedRefRumpunJabatan(offset int, limit int) ([]migrations.RefRumpunJabatan, error)
	Store(rumpunJabatan migrations.RefRumpunJabatan) (migrations.RefRumpunJabatan, error)
	FindLast() (migrations.RefRumpunJabatan, error)
	FindById(ID string) (migrations.RefRumpunJabatan, error)
	CountActiveRefRumpunJabatan() int64
}

type RefSubRumpunJabatanRepository interface {
	FindAllRefSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error)
	FindAllActiveRefSubRumpunJabatan() ([]migrations.RefSubrumpunJabatan, error)
	FindAllPaginatedRefSubRumpunJabatan(offset int, limit int) ([]migrations.RefSubrumpunJabatan, error)
	CountActiveRefSubRumpunJabatan() int64
}
