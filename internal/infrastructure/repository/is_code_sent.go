package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) IsCodeSent(c context.Context, phone string) (bool, error) {
	var code string
	sql := "SELECT code FROM confirmation_codes WHERE phone = $1 AND created_at + INTERVAL '3 minutes' > NOW();"
	err := r.db.QueryRow(c, sql, phone).Scan(&code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
