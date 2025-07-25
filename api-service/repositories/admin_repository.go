package repositories

import (
	"api-service/models"
	"api-service/utils"
	"database/sql"
)

type AdminRepository interface {
	Insert(admin *models.Admin) error
	QueryById(id int) (*models.Admin, error)
}

type adminRepository struct {
	Db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{Db: db}
}

func (repo *adminRepository) Insert(admin *models.Admin) error {
	query := `
	INSERT INTO Admins(Name, Password, Email)
	OUTPUT INSERTED.AdminId
	VALUES(@name, @password, @email)`

	statement, err := repo.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return err
	}

	var adminId int
	err = statement.QueryRow(
		sql.Named("name", admin.Name),
		sql.Named("password", hashedPassword),
		sql.Named("email", admin.Email)).Scan(&adminId)
	if err != nil {
		return err
	}
	admin.AdminId = adminId
	return nil
}

func (repo *adminRepository) QueryById(id int) (*models.Admin, error) {
	query := `SELECT AdminId, Name, Password, Email FROM Admins WHERE AdminId = @adminId`
	row := repo.Db.QueryRow(query, sql.Named("adminId", id))
	var admin models.Admin
	err := row.Scan(&admin.AdminId, &admin.Name, &admin.Password, &admin.Email)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
