package controllers

import (
	"fmt"
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
		"message": "success fetch data",
		"data": users,
	})
}

func (r *UserController) Detail(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var user models.User
	err := facades.Orm().Query().Select("id", "name", "email").Where("id", id).First(&user)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}
	return ctx.Response().Json(200, map[string]any {
		"success": true,
		"message": "success fetch data",
		"data": user,
	})
}

func (r *UserController) Create(ctx http.Context) http.Response {
	var user models.User
	err := ctx.Request().Bind(&user)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}
	password := user.Password
	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		fmt.Println(err.Error())
	}
	user.Password = hashedPassword
	err = facades.Orm().Query().Create(&user)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}
	return ctx.Response().Json(200, map[string]any {
		"success": true,
		"message": "success create data",
	})
}

func (r *UserController) Update(ctx http.Context) http.Response {
	var user models.User
	id := ctx.Request().Route("id")
	err := ctx.Request().Bind(&user)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}
	password := user.Password
	hashedPassword, err := facades.Hash().Make(password)
	user.Password = hashedPassword
	_, err = facades.Orm().Query().Where("id", id).Update(&user)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}
	return ctx.Response().Json(200, map[string]any {
		"success": true,
		"message": "success update data",
	})
}

func (r *UserController) Delete(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var user models.User
	_, err := facades.Orm().Query().Where("id", id).Delete(&user)
	if err != nil {
		return ctx.Response().Json(500, map[string]any {
			"success": false,
			"message": err.Error(),
		})
	}
	return ctx.Response().Json(200, map[string]any{
		"success": false,
		"message": "success delete data",
	})
}