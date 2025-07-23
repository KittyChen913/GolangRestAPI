package models

import (
	"time"
)

type User struct {
	Id             int
	Name           string `binding:"required"`
	Age            int
	CreateDateTime time.Time
}
