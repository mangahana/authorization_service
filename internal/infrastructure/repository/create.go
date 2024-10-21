package repository

import (
	"authorization_service/internal/core/dto"
	"context"
)

func (r *repo) Create(c context.Context, dto *dto.Register) (int, error) {
	var id int
	sql := "INSERT INTO users (phone, username, password) VALUES ($1, $2, $3) RETURNING id;"
	err := r.db.QueryRow(c, sql, dto.Phone, dto.Username, dto.Password).Scan(&id)
	return id, err
}
