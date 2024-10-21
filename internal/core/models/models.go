package models

type User struct {
	ID          int     `json:"id"`
	Username    string  `json:"username"`
	Description string  `json:"description"`
	Photo       *string `json:"photo"`
	IsBanned    bool    `json:"is_banned"`
}

type LoginCredentials struct {
	UserID   int
	Password string
}
