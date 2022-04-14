package datausecases

import (
	dataprotocols "sales-tool-clone/src/data/protocols"
	domain_usecases "sales-tool-clone/src/domain/usecases"
)

type DbAddUser struct {
	UserRepository dataprotocols.UserRepository
}

func (d DbAddUser) Add(account domain_usecases.AddUserParams) bool {
	success := d.UserRepository.Add(account)
	return success
}
