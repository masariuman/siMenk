package migrations

type RefSubrumpunJabatan struct {
	Id_subrumpun     int    `gorm:"primaryKey; type:int(11)" json:"id_subrumpun"`
	Id_rumpun        int    `gorm:"type:int(11)" json:"id_rumpun"`
	Subrumpun        string `gorm:"type:varchar(500)" json:"subrumpun"`
	Definisi         string `gorm:"type:text" json:"definisi"`
	Perangkat_daerah string `gorm:"type:text" json:"perangkat_daerah"`
}
