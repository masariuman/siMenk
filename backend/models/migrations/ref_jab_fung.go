package migrations

type RefJabFung struct {
	kode             string `gorm:"type:carchar(10)" "primaryKey" json:"kode"`
	kode_jns         string `gorm:"type:varchar(1)" json:"kode_jns"`
	nama_jns         string `gorm:"type:varchar(50)" json:"nama_jns"`
	kode_rumpun      string `gorm:"type:varchar(2)" json:"kode_rumpun"`
	nama_rumpun      string `gorm:"type:varchar(100)" json:"nama_rumpun"`
	kode_kategori    string `gorm:"type:varchar(2)" json:"kode_kategori"`
	nama_kategori    string `gorm:"type:varchar(50)" json:"nama_kategori"`
	kode_jab         string `gorm:"type:varchar(3)" json:"kode_jab"`
	nama_jab         string `gorm:"type:varchar(100)" json:"nama_jab"`
	kode_jenjang     string `gorm:"type:varchar(2)" json:"kode_jenjang"`
	nama_jenjang     string `gorm:"type:varchar(50)" json:"nama_jenjang"`
	nama_jab_singkat string `gorm:"type:varchar(50)" json:"nama_jab_singkat"`
	nama_jab_lama    string `gorm:"type:varchar(100)" json:"nama_jab_lama"`
	kualifikasi      string `gorm:"type:text" json:"kualifikasi"`
	tugas            string `gorm:"type:text" json:"tugas"`
	pembina          string `gorm:"type:varchar(100)" json:"pembina"`
	aturan           string `gorm:"type:text" json:"aturan"`
	kelas_jab        string `gorm:"type:varchar(2)" json:"kelas_jab"`
	bup              int    `gorm:"type:int(11)" json:"bup"`
	aktif            string `gorm:"type:varchar(1)" json:"aktif"`
	tunjangan_jab    int    `gorm:"type:double" json:"tunjangan_jab"`
	is_deleted       int    `gorm:"type:int(11)" json:"is_deleted"`
}
