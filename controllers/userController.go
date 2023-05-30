package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/helpers"
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/models"
	userResponse "task-5-vix-btpns-Moh.AinurBahtiarRohman/app"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

func (h *userController) Register(c *echo.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	user.Password = helpers.HashPassword(user.Password)

	err := h.db.Debug().Create(&user).Error
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := echo.H{
			"errors": errors,
		}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userResponse.FormatUserResponse(user, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User Registered Succesfully")

	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *echo.Context) {
	var user models.User

	c.ShouldBindJSON(&user)

	Inputpassword := user.Password
	err := h.db.Debug().Where("email = ?", user.Email).Find(&user).Error
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

	c.JSON(http.StatusOK, response)
}

func (h *userController) Update(c *echo.Context) {
	var oldUser models.User
	var newUser models.User

	id := c.Param("userId")

	err := h.db.First(&oldUser, id).Error
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := echo.H{
			"errors": errors,
		}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newUser)
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
	c.JSON(http.StatusOK, response)
}

func (h *userController) Delete(c *echo.Context) {
	var user models.User

	id := c.Param("userId")
	err := h.db.First(&user, id).Error
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
	c.JSON(http.StatusOK, response)
}
