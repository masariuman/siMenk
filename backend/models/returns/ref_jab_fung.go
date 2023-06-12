package returns

type RefJabFung struct {
	Kode         string `gorm:"primaryKey; type:varchar(10)" json:"kode"`
	Nama_jab     string `gorm:"type:varchar(100)" json:"nama_jab"`
	Nama_jenjang string `gorm:"type:varchar(50)" json:"nama_jenjang"`
}

type Tabler interface {
	TableName() string
}

func (RefJabFung) TableName() string {
	return "ref_jab_fung"
}
