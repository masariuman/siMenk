package migrations

type Tabler interface {
	TableName() string
}

func (RefJabFung) TableName() string {
	return "ref_jab_fung"
}

func (RefRumpunJabatan) TableName() string {
	return "ref_jab_fung"
}

func (RefSubrumpunJabatan) TableName() string {
	return "ref_jab_fung"
}
