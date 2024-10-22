package repository

import "context"

func (r *repo) GetUserIDbyUsername(c context.Context, username string) (int, error) {
	var userId int
	sql := "SELECT id FROM users WHERE username = $1;"
	return userId, r.db.QueryRow(c, sql, username).Scan(&userId)
}
