package controllers

import (
	"encoding/json"
	"net/http"

	userResponse "github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/app"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/helpers"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

func (h *userController) Register(c echo.Context) (err error) {
	var user models.User
	c.Bind(&user)

	user.Password = helpers.HashPassword(user.Password)

	err = h.db.Debug().Create(&user).Error
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := echo.Map{
			"errors": errors,
		}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userResponse.FormatUserResponse(user, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User Registered Succesfully")

	return c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c echo.Context) (err error) {
	var user models.User

	c.Bind(&user)

	Inputpassword := user.Password
	err = h.db.Debug().Where("email = ?", user.Email).Find(&user).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	comparePass := helpers.ComparePassword(user.Password, Inputpassword)
	if !comparePass {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userResponse.FormatUserResponse(user, token)
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User Login Succesfully")

	return c.JSON(http.StatusOK, response)
}

func (h *userController) Update(c echo.Context) (err error) {
	var oldUser models.User
	var newUser models.User

	id := c.Param("userId")

	err = h.db.First(&oldUser, id).Error
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := echo.Map{
			"errors": errors,
		}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.db.Model(&oldUser).Updates(newUser).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponse(http.StatusOK, "success", nil, "User Updated Succesfully")
	return c.JSON(http.StatusOK, response)
}

func (h *userController) Delete(c echo.Context) (err error) {
	var user models.User

	id := c.Param("userId")
	err = h.db.First(&user, id).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.db.Delete(&user).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponse(http.StatusOK, "success", nil, "User Deleted Succesfully")
	return c.JSON(http.StatusOK, response)
}
