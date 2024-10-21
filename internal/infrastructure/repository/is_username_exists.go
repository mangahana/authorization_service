package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) IsUsernameExists(c context.Context, username string) (bool, error) {
	var id int
	sql := "SELECT id FROM users WHERE username = $1;"
	err := r.db.QueryRow(c, sql, username).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
