package dataprotocols

import (
	"sales-tool-clone/src/domain/usecases"
)


type UserRepository interface {
	Add(userData domain_usecases.AddUserParams) bool
}