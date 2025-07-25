package services

import (
	"api-service/models"
	"api-service/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetUser(userId int) (*models.User, error)
	UpdateUser(userId int, user *models.User) error
	DeleteUser(userId int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) CreateUser(user *models.User) error {
	err := u.repo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetUsers() ([]*models.User, error) {
	users, err := u.repo.Query()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) GetUser(userId int) (*models.User, error) {
	user, err := u.repo.QueryById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) UpdateUser(userId int, user *models.User) error {
	_, err := u.repo.QueryById(int(userId))
	if err != nil {
		return err
	}
	user.Id = int(userId)
	err = u.repo.Update(user)
	if err != nil {

		return err
	}
	return nil
}

func (u *userService) DeleteUser(userId int) error {
	user, err := u.repo.QueryById(userId)
	if err != nil {
		return err
	}
	err = u.repo.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
