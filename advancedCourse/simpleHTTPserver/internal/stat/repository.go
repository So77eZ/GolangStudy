package stat

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
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
			Clics:  1,
			Date:   currentDate,
		})
	} else {
		stat.Clics += 1
		repo.Db.Save(&stat)
	}
}
