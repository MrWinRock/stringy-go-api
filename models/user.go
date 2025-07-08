package models

type User struct {
	UserID            int     `db:"user_id" json:"user_id"`
	Username          string  `db:"username" json:"username"`
	Email             string  `db:"email" json:"email"`
	Role              string  `db:"role" json:"role"`
	CreatedAt         string  `db:"created_at" json:"created_at"`
	ProfilePictureURL *string `db:"profile_picture_url" json:"profile_picture_url"`
	Bio               *string `db:"bio" json:"bio,omitempty"`
}
