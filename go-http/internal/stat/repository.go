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

func (repo *StatRepository) GetStat(from, to time.Time, by string) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	repo.Db.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Scan(&stats)
	return stats
}
