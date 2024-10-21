package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateSession(c context.Context, userId int) (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	sql := "INSERT INTO sessions (user_id, token) VALUES ($1, $2);"
	err = r.db.QueryRow(c, sql, userId, token).Scan()
	if err == nil || errors.Is(err, pgx.ErrNoRows) {
		return token.String(), nil
	}
	return "", err
}
