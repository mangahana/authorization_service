package repository

import (
	"authorization_service/internal/core/models"
	"context"
)

func (r *repo) GetByToken(c context.Context, token string) (models.UserSession, error) {
	var u models.UserSession
	u.Permissions = []string{}
	sql := `SELECT id, username, photo, is_banned,
						(SELECT COALESCE(array_agg(permissions.name), '{}') FROM roles
							RIGHT JOIN permissions ON permissions.id = any(roles.permissions)
						WHERE roles.id = role_id) as permissions
					FROM users WHERE id = (SELECT user_id FROM sessions WHERE token = $1)`
	err := r.db.QueryRow(c, sql, token).Scan(&u.ID, &u.Username, &u.Photo, &u.IsBanned, &u.Permissions)
	return u, err
}
