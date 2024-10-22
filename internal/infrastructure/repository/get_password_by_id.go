package repository

import "context"

func (r *repo) GetPasswordByID(c context.Context, id int) (string, error) {
	var password string
	sql := "SELECT password FROM users WHERE id = $1;"
	return password, r.db.QueryRow(c, sql, id).Scan(&password)
}
