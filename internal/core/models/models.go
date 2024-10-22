package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	IsBanned    bool   `json:"is_banned"`
}

type UserSession struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Photo       string   `json:"photo"`
	IsBanned    bool     `json:"is_banned"`
	Permissions []string `json:"permissions"`
}

type LoginCredentials struct {
	UserID   int
	Password string
}
