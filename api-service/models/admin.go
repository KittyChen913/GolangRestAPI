package models

import (
	"api-service/db"
	"api-service/utils"
	"database/sql"
)

type Admin struct {
	AdminId  int
	Name     string `binding:"required"`
	Password string `binding:"required"`
	Email    string
}

func (a *Admin) Insert() error {
	query := `
	INSERT INTO Admins(Name, Password, Email)
	OUTPUT INSERTED.AdminId
	VALUES(@name, @password, @email)`

	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(a.Password)
	if err != nil {
		return err
	}

	var adminId int
	err = statement.QueryRow(
		sql.Named("name", a.Name),
		sql.Named("password", hashedPassword),
		sql.Named("email", a.Email)).Scan(&adminId)
	if err != nil {
		return err
	}
	a.AdminId = adminId
	return nil
}
