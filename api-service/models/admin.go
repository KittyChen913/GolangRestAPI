package models

type Admin struct {
	AdminId  int
	Name     string `binding:"required"`
	Password string `binding:"required"`
	Email    string
}
