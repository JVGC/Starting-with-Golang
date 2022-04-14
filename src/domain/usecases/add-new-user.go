package domain_usecases

import (
	"sales-tool-clone/src/domain/models"
)

type AddUserParams struct {
	Name string
	Email string
	Password string
	Role models.UserRoles
}

type AddUser interface {
	Add(account AddUserParams) bool
}