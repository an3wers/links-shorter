package user

import (
	"go/links-shorter/pkg/db"
	"log"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Db *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{Db: db}
}

func (repository *UserRepository) GetByEmail(email string) (*User, error) {
	var user User

	result := repository.Db.First(&user, "email = ?", email)

	log.Println("result: ", result)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepository) GetById(id uint) (*User, error) {
	var user User

	result := repository.Db.First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepository) Create(user *User) (*User, error) {
	result := repository.Db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UserRepository) Update(user *User) (*User, error) {
	result := repository.Db.Clauses(clause.Returning{}).Updates(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UserRepository) DeleteById(id uint) error {
	result := repository.Db.Delete(&User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
