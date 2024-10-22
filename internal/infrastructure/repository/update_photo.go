package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) UpdatePhoto(c context.Context, userId int, filename string) error {
	sql := "UPDATE users SET photo = $2 WHERE id = $1;"
	err := r.db.QueryRow(c, sql, userId, filename).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
