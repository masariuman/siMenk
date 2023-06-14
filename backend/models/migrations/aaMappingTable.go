package migrations

type Tabler interface {
	TableName() string
}

func (RefJabFung) TableName() string {
	return "ref_jab_fung"
}

func (RefRumpunJabatan) TableName() string {
	return "rumpun_jabatan"
}

func (RefSubrumpunJabatan) TableName() string {
	return "subrumpun_jabatan"
}
