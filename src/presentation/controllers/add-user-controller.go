package controllers

import (
	"fmt"
	"net/http"
	"sales-tool-clone/src/domain/models"
	domain_usecases "sales-tool-clone/src/domain/usecases"

	"github.com/gin-gonic/gin"
)

type AddUserController struct {
	AddUser domain_usecases.AddUser
}

type AddUserRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     models.UserRoles `json:"role"`
}

func (d AddUserController) Handle(c *gin.Context) {

	var requestBody AddUserRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	fmt.Println(requestBody)

	success := d.AddUser.Add(domain_usecases.AddUserParams{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
		Role:     requestBody.Role,
	})

	c.IndentedJSON(http.StatusOK, gin.H{"success": success})
}
