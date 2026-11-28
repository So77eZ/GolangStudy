package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Поддерживаемые значения параметра by в GET /stat.
const (
	PeriodForMonth string = "month"
	PeriodForDay   string = "day"
)

// Stat хранит число переходов по ссылке за один календарный день.
type Stat struct {
	gorm.Model
	LinkId uint           `json:"link_id"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date"`
}
