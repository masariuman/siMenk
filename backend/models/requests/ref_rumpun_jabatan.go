package requests

type RefRumpunJabatanRequest struct {
	Rumpun string `gorm:"type:varchar(255)" json:"rumpun" binding:"required"`
}
