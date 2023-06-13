package returns

type RefRumpunJabatan struct {
	Id_rumpun int    `gorm:"primaryKey" json:"id_rumpun"`
	Nomor     int    `json:"nomor"`
	Rumpun    string `gorm:"type:varchar(255)" json:"rumpun"`
}
