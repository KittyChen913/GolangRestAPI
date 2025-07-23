package repositories

import (
	"api-service/db"
	"api-service/models"
	"database/sql"
)

type UserRepository interface {
	Insert(user *models.User) error
	Query() ([]*models.User, error)
	QueryById(id int) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}

type userRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{Db: db}
}

func (repo *userRepository) Insert(user *models.User) error {
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
		sql.Named("name", user.Name),
		sql.Named("age", user.Age),
		sql.Named("createDateTime", user.CreateDateTime)).Scan(&userId)
	if err != nil {
		return err
	}
	user.Id = userId
	return nil
}

func (repo *userRepository) Query() ([]*models.User, error) {
	query := `SELECT Id, Name, Age, CreateDateTime FROM Users`
	result, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var users []*models.User

	for result.Next() {
		var user models.User
		err := result.Scan(&user.Id, &user.Name, &user.Age, &user.CreateDateTime)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repo *userRepository) QueryById(id int) (*models.User, error) {
	query := `SELECT Id, Name, Age, CreateDateTime FROM Users WHERE Id = @userId`
	result := db.Db.QueryRow(query, sql.Named("userId", id))
	var user models.User
	err := result.Scan(&user.Id, &user.Name, &user.Age, &user.CreateDateTime)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Update(user *models.User) error {
	query := `
	UPDATE Users
	SET Name = @name, Age = @age
	WHERE Id = @userId`

	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		sql.Named("name", user.Name),
		sql.Named("age", user.Age),
		sql.Named("userId", user.Id))
	return err
}

func (repo *userRepository) Delete(user *models.User) error {
	query := `
	DELETE 
	FROM Users
	WHERE Id = @userId`

	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(sql.Named("userId", user.Id))
	return err
}
