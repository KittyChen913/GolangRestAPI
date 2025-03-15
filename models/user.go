package models

import (
	"database/sql"
	"golangrestapi/db"
	"time"
)

type User struct {
	Id             int
	Name           string `binding:"required"`
	Age            int
	CreateDateTime time.Time
}

func (u *User) Insert() error {
	query := `
	INSERT INTO Users(Name, Age, CreateDateTime)
	OUTPUT INSERTED.ID
	VALUES(@name, @age, @createDateTime)`

	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	var userId int
	err = statement.QueryRow(
		sql.Named("name", u.Name),
		sql.Named("age", u.Age),
		sql.Named("createDateTime", u.CreateDateTime)).Scan(&userId)
	if err != nil {
		return err
	}
	u.Id = userId
	return nil
}
