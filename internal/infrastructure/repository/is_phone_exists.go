package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) IsPhoneExists(c context.Context, phone string) (bool, error) {
	var id int
	sql := "SELECT id FROM users WHERE phone = $1;"

	err := r.db.QueryRow(c, sql, phone).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
