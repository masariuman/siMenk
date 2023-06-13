package returns

type RefSubrumpunJabatan struct {
	Id_subrumpun     int    `gorm:"primaryKey" json:"id_subrumpun"`
	Nomor            int    `json:"nomor"`
	Id_rumpun        int    `json:"id_rumpun"`
	Subrumpun        string `gorm:"type:varchar(500)" json:"subrumpun"`
	Definisi         string `gorm:"type:text" json:"definisi"`
	Perangkat_daerah string `gorm:"type:text" json:"perangkat_daerah"`
}
