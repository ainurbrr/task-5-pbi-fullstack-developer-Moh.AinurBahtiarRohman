package controllers

import (
	"io"
	"net/http"
	"os"

	photoResponse "github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/app"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/helpers"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/models"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type photoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *photoController {
	return &photoController{db}
}

func (h *photoController) Get(c echo.Context) (err error) {
	var userPhoto models.Photo
	err = h.db.Preload("User").Find(&userPhoto).Error

	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Failed to Get Your Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userPhoto.PhotoURL == "" {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Please Upload Your Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photoResponse.FormatPhotoResponse(&userPhoto)
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "Successfully Fetch User Photo")
	return c.JSON(http.StatusOK, response)
}

func (h *photoController) Create(c echo.Context) (err error) {
	var userPhoto models.Photo
	var countPhoto int64
	currentUser := c.Get("currentUser").(models.User)

	h.db.Model(&userPhoto).Where("user_id = ?", currentUser.ID).Count(&countPhoto)
	if countPhoto > 0 {
		data := echo.Map{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "You Already Have a Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err = c.Bind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := echo.Map{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Failed to Upload User Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("photo_profile")
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := echo.Map{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Failed to Upload User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "images/avatar/" + uuid.New().String() + extension

	//upload the avatar
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
	// Create a new file on disk
	dst, err := os.Create(path)
	if err != nil {
		return
	}
	defer dst.Close()
	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	if err != nil {
		data := echo.Map{"is_uploaded": false}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Failed to Upload User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	h.InsertPhoto(input, path, currentUser.ID)

	data := echo.Map{"is_uploaded": true}
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Uploaded")
	return c.JSON(http.StatusOK, response)
}

func (h *photoController) InsertPhoto(userPhoto models.Photo, fileLocation string, currUserID uint) error {
	savePhoto := models.Photo{
		UserID:   currUserID,
		Title:    userPhoto.Title,
		Caption:  userPhoto.Caption,
		PhotoURL: fileLocation,
	}

	err := h.db.Debug().Create(&savePhoto).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *photoController) Update(c echo.Context) (err error) {
	var userPhoto models.Photo
	currentUser := c.Get("currentUser").(models.User)

	err = h.db.Where("user_id = ?", currentUser.ID).Find(&userPhoto).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Update")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input = userPhoto
	err = c.Bind(&input)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Update")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("update_profile")
	if err != nil {
		data := echo.Map{"is_uploaded": false}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Failed to Update User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "images/avatar/" + uuid.New().String() + extension

	//upload the avatar
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
	// Create a new file on disk
	dst, err := os.Create(path)
	if err != nil {
		return
	}
	defer dst.Close()
	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Upload")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.Bind(userPhoto)
	userPhoto.UserID = currentUser.ID
	userPhoto.User = &currentUser

	h.UpdatePhoto(input, &userPhoto, path)

	data := photoResponse.FormatPhotoResponse(&userPhoto)
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Updated")
	return c.JSON(http.StatusOK, response)
}

func (h *photoController) UpdatePhoto(oldPhoto models.Photo, newPhoto *models.Photo, path string) error {
	newPhoto.Title = oldPhoto.Title
	newPhoto.Caption = oldPhoto.Caption
	newPhoto.PhotoURL = path

	err := h.db.Updates(&newPhoto).Error
	if err != nil {
		return err
	}

	return nil
}

func (h *photoController) Delete(c echo.Context) (err error) {
	var userPhoto models.Photo
	currentUser := c.Get("currentUser").(models.User)

	err = h.db.Where("user_id = ?", currentUser.ID).Delete(&userPhoto).Error
	if err != nil {
		data := echo.Map{
			"is_deleted": false,
		}

		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "Failed to delete user photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := echo.Map{
		"is_deleted": true,
	}

	response := helpers.ApiResponse(http.StatusOK, "success", data, "User Photo Successfully Deleted")
	return c.JSON(http.StatusOK, response)
}
