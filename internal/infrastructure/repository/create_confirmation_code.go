package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateConfirmationCode(c context.Context, phone, code, ip string) error {
	sql := "INSERT INTO confirmation_codes (phone, code, ip) VALUES($1, $2, $3);"
	err := r.db.QueryRow(c, sql, phone, code, ip).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
