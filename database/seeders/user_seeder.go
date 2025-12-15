package seeders

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UserSeeder struct {
	Name     string
	Email    string
	Password string
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	password := "secret"
	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		fmt.Println("Error hashing", err.Error())
	}
	user := models.User{
		Name:     "atmin",
		Email:    "admin@gmail.com",
		Password: hashedPassword,
	}
	return facades.Orm().Query().Create(&user)
}
