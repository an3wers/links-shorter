package link

import "go/links-shorter/pkg/db"

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
