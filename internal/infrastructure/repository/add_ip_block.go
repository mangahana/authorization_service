package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) AddIPblock(c context.Context, ip string) error {
	sql := "INSERT INTO ip_block_list (ip) VALUES ($1);"
	err := r.db.QueryRow(c, sql, ip).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
