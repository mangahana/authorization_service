package repository

import (
	"authorization_service/internal/core/models"
	"context"
)

func (r *repo) GetCredentialsByPhone(c context.Context, phone string) (models.LoginCredentials, error) {
	var output models.LoginCredentials
	sql := "SELECT id, password FROM users WHERE phone = $1;"
	err := r.db.QueryRow(c, sql, phone).Scan(&output.UserID, &output.Password)
	return output, err
}
