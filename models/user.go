package models

type User struct {
	UserID            int     `db:"user_id" json:"user_id"`
	Username          string  `db:"username" json:"username"`
	Email             string  `db:"email" json:"email"`
	Password          string  `db:"password" json:"password"`
	Role              string  `db:"role" json:"role"`
	CreatedAt         string  `db:"created_at" json:"created_at"`
	ProfilePictureURL *string `db:"profile_picture_url" json:"profile_picture_url"`
}

type UserLogin struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
