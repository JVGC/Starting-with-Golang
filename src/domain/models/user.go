package models

type UserRoles string

const  (
	Admin UserRoles = "admin"
	Salesperson = "salesperson"
	Consultant = "consultant"
)

type UserModel struct {
	Id string
	Name string
	Email string
	Password string
	Role UserRoles
}

