package models

type Room struct {
	RoomID         string  `db:"room_id" json:"room_id"`
	Title          string  `db:"title" json:"title"`
	Description    string  `db:"description" json:"description"`
	RoomPictureURL *string `db:"room_picture_url" json:"room_picture_url"` // optional
}
