package stat

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Db *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

// AddClick создаёт дневной счётчик при первом переходе или увеличивает его.
func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	// Ищем счётчик по ID ссылки и текущей дате.
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

// GetBy суммирует клики по периодам в диапазоне с включёнными границами.
func (repo *StatRepository) GetBy(from, to datatypes.Date, period string) ([]StatPeriod, error) {
	// Пустой результат кодируется как JSON-массив [].
	stats := make([]StatPeriod, 0)

	// Начало дня или месяца служит ключом группировки для SUM.
	err := repo.Db.
		Table("stats").
		Select("date_trunc(?,date)::date AS period, SUM(clicks) as clicks", period).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats).Error
	if err != nil {
		return nil, err
	}
	return stats, nil
}
