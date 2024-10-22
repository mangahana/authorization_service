package repository

import (
	"authorization_service/internal/core/dto"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) UpdateUser(c context.Context, userId int, dto *dto.Update) error {
	sql := `UPDATE users SET username = $1, description = $2 WHERE id = $3;`

	err := r.db.QueryRow(c, sql, dto.Username, dto.Description, userId).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
