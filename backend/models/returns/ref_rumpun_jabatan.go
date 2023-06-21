package returns

type RefRumpunJabatan struct {
	Id_rumpun int    `gorm:"primaryKey" json:"id_rumpun"`
	Nomor     string `json:"nomor"`
	Rumpun    string `gorm:"type:varchar(255)" json:"rumpun"`
}
