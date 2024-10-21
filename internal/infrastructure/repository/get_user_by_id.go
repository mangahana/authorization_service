package repository

import (
	"authorization_service/internal/core/models"
	"context"
)

func (r *repo) GetUserByID(c context.Context, id int) (models.User, error) {
	var u models.User
	sql := "SELECT id, username, description, photo, is_banned FROM users WHERE id = $1;"
	return u, r.db.QueryRow(c, sql, id).Scan(&u.ID, &u.Username, &u.Description, &u.Photo, &u.IsBanned)
}
