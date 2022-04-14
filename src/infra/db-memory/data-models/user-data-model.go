package datamodels

import (
	"sales-tool-clone/src/domain/models"
	"time"
)

type UserDataModel struct {
	Id        string
	Name      string
	Email     string
	Password  string
	Role      models.UserRoles
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
