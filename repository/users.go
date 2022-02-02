package repository

import (
	"errors"
	"restsample/db"
	"restsample/db/models"
)

type userRepo struct {
	db *db.UserDb
}

type UserRepo interface {
	FindOne(query models.UserQuery) (*models.User, error)
	FindAll(query models.UserQuery) ([]models.User, error)
}

// find user by id
func (repo *userRepo) FindOne(query models.UserQuery) (*models.User, error) {
	if query.Id == nil {
		return nil, errors.New("no id present for query in FindOne")
	}
	for _, x := range repo.db.Users {
		if x.Id == *query.Id {
			return &x, nil
		}
	}
	return nil, nil
}

// find all user by ids
func (repo *userRepo) FindAll(query models.UserQuery) ([]models.User, error) {
	var res []models.User
	for _, x := range repo.db.Users {
		_, exist := query.IdsMap[x.Id]
		if exist {
			res = append(res, x)
		}
	}
	return res, nil
}
func NewUserRepo(db *db.UserDb) UserRepo {
	return &userRepo{
		db: db,
	}
}
