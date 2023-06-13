package responses

type RefJabFung struct {
	Kode             string `gorm:"primaryKey; type:varchar(10)" json:"kode"`
	Kode_jns         string `gorm:"type:varchar(1)" json:"kode_jns"`
	Nama_jns         string `gorm:"type:varchar(50)" json:"nama_jns"`
	Kode_rumpun      string `gorm:"type:varchar(2)" json:"kode_rumpun"`
	Nama_rumpun      string `gorm:"type:varchar(100)" json:"nama_rumpun"`
	Kode_kategori    string `gorm:"type:varchar(2)" json:"kode_kategori"`
	Nama_kategori    string `gorm:"type:varchar(50)" json:"nama_kategori"`
	Kode_jab         string `gorm:"type:varchar(3)" json:"kode_jab"`
	Nama_jab         string `gorm:"type:varchar(100)" json:"nama_jab"`
	Kode_jenjang     string `gorm:"type:varchar(2)" json:"kode_jenjang"`
	Nama_jenjang     string `gorm:"type:varchar(50)" json:"nama_jenjang"`
	Nama_jab_singkat string `gorm:"type:varchar(50)" json:"nama_jab_singkat"`
	Nama_jab_lama    string `gorm:"type:varchar(100)" json:"nama_jab_lama"`
	Kualifikasi      string `gorm:"type:text" json:"kualifikasi"`
	Tugas            string `gorm:"type:text" json:"tugas"`
	Pembina          string `gorm:"type:varchar(100)" json:"pembina"`
	Aturan           string `gorm:"type:text" json:"aturan"`
	Kelas_jab        string `gorm:"type:varchar(2)" json:"kelas_jab"`
	Bup              int    `gorm:"type:int(11)" json:"bup"`
	Aktif            string `gorm:"type:varchar(1)" json:"aktif"`
	Tunjangan_jab    int    `gorm:"type:double" json:"tunjangan_jab"`
	Is_deleted       int    `gorm:"type:int(11)" json:"is_deleted"`
}
