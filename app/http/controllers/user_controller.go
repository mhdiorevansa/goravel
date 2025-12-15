package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	var users []models.User

	err := facades.Orm().Query().Select("id", "name", "email").Get(&users)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.Response().Json(200, map[string]any {
		"success": true,
		"data": users,
	})
}
