package repositories

import (
	"fmt"
	domain_usecases "sales-tool-clone/src/domain/usecases"
	datamodels "sales-tool-clone/src/infra/db-memory/data-models"
	"time"
)

type MemoryUserRepository struct {
	Users []datamodels.UserDataModel
}

func (d *MemoryUserRepository) Add(userData domain_usecases.AddUserParams) bool {
	newUser := datamodels.UserDataModel{
		Id:        "husauhsauhashu",
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Role:      userData.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Date(2023, time.August, time.Now().Day(), 0, 0, 0, 0, time.Local),
	}
	d.Users = append(d.Users, newUser)
	fmt.Println(len(d.Users))
	return true
}
