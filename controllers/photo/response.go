package photo

import (
	"time"
)

type photo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type createPhoto struct {
	photo
	CreatedAt time.Time `json:"created_at"`
}

type updatePhoto struct {
	photo
	UpdatedAt time.Time `json:"updated_at"`
}
