package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/facades"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users", userController.Show)
}
