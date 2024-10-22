package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) UpdatePassword(c context.Context, userId int, password string) error {
	sql := "UPDATE users SET password = $2 WHERE id = $1;"
	err := r.db.QueryRow(c, sql, userId, password).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
