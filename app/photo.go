package app

import "github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/models"

type PhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
	User     models.User
}

func FormatPhotoResponse(photo *models.Photo) interface{} {
	var formatter interface{}

	formatter = PhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.User.ID,
		User:     *photo.User,
	}

	return formatter
}
