package stat

import (
	"errors"
	"go/links-shorter/pkg/db"
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

func (repository *StatRepository) AddClick(linkId uint) error {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repository.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)

	if stat.ID == 0 {
		// create new stat
		repository.Db.Create(&Stat{
			LinkId: linkId,
			Date:   currentDate,
			Clicks: 1,
		})

	} else {
		// update existing stat
		stat.Clicks++
		repository.Db.Save(&stat)
	}

	return nil
}

func (repository *StatRepository) GetStat(from, to time.Time, by string) ([]GetStatSqlResponse, error) {
	var stats []GetStatSqlResponse
	var selectQuery string

	switch by {
	case "day":
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks) as sum"
	case "month":
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks) as sum"
	case "year":
		selectQuery = "to_char(date, 'YYYY') as period, sum(clicks) as sum"
	default:
		return nil, errors.New("invalid by")
	}

	repository.Db.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats, nil
}
