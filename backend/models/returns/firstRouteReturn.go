package returns

import "time"

type Test struct {
	Id        uint      `gorm:"primaryKey"`
	Test      string    `gorm:"type:varchar(255)" json:"test"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}
