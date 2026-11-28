package stat

import (
	"gorm.io/datatypes"
)

// GetStatResponse содержит список периодов с суммами кликов.
type GetStatResponse struct {
	Stats []StatPeriod `json:"stat"`
}

// StatPeriod содержит дату начала периода и число кликов за него.
type StatPeriod struct {
	Period datatypes.Date `json:"period"`
	Clicks int64          `json:"clicks"`
}
