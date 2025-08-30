package stat

import (
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
