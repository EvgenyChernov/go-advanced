package stat

import (
	"app/adv-http/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{Db: db}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	var currentDate = datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.Id == 0 {
		repo.Db.Create(&Stat{
			LinkID: linkId,
			Date:   currentDate,
			Clicks: 1,
		})
	} else {
		stat.Clicks++
		repo.Db.Save(&stat)
	}
}
