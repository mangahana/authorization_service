package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) IsPhoneBlocked(c context.Context, phone string) (bool, error) {
	var output string
	sql := "SELECT phone FROM phone_block_list WHERE phone = $1;"
	err := r.db.QueryRow(c, sql, phone).Scan(&output)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
