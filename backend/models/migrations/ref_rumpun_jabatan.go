package migrations

type RefRumpunJabatan struct {
	Id_rumpun int    `gorm:"primaryKey; type:int(11)" json:"id_rumpun"`
	Rumpun    string `gorm:"type:varchar(255)" json:"rumpun"`
}
