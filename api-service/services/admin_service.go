package services

import (
	"api-service/models"
	"api-service/repositories"
)

type AdminService interface {
	SignUpAdmin(admin *models.Admin) error
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (a *adminService) SignUpAdmin(admin *models.Admin) error {
	err := a.repo.Insert(admin)
	if err != nil {
		return err
	}
	return nil
}
