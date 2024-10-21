package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) IsIpBlocked(c context.Context, ip string) (bool, error) {
	var output string
	sql := "SELECT ip FROM ip_block_list WHERE ip = $1;"
	err := r.db.QueryRow(c, sql, ip).Scan(&output)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
