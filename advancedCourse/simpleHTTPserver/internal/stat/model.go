package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Stat хранит число переходов по ссылке за один календарный день.
type Stat struct {
	gorm.Model
	LinkId uint           `json:"link_id"`
	Clics  int            `json:"clics"`
	Date   datatypes.Date `json:"date"`
}
