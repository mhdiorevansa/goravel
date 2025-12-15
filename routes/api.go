package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/facades"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users", userController.Show)
	facades.Route().Prefix("user").Get("/{id}", userController.Detail)
	facades.Route().Prefix("user").Post("/create", userController.Create)
	facades.Route().Prefix("user").Put("/update/{id}", userController.Update)
	facades.Route().Prefix("user").Delete("/delete/{id}", userController.Delete)
}
