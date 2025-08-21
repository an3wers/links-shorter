package link

import (
	"go/links-shorter/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Db *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{Db: db}
}

func (repository *LinkRepository) CreateLink(link *Link) (*Link, error) {
	result := repository.Db.Create(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repository *LinkRepository) GetByHash(hash string) (*Link, error) {
	var found Link
	result := repository.Db.First(&found, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &found, nil
}

func (repository *LinkRepository) UpdateLink(link *Link) (*Link, error) {
	result := repository.Db.Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repository *LinkRepository) DeleteById(id uint) error {
	result := repository.Db.Delete(&Link{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *LinkRepository) GetLinkById(id uint) (*Link, error) {
	var link Link
	result := repository.Db.First(&link, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
