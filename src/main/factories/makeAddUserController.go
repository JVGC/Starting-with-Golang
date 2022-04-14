package factories

import (
	datausecases "sales-tool-clone/src/data/usecases"
	"sales-tool-clone/src/infra/db-memory/repositories"
	"sales-tool-clone/src/presentation"
	"sales-tool-clone/src/presentation/controllers"
)

func MakeAddUserController() presentation.Controller {

	memoryUserRepository := repositories.MemoryUserRepository{}
	dbAddUser := datausecases.DbAddUser{UserRepository: &memoryUserRepository}

	addUserController := controllers.AddUserController{AddUser: dbAddUser}

	return addUserController

}
