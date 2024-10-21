package repository

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) GetSMSByCredentials(c context.Context, dto *dto.ConfirmPhone) error {
	var code string
	sql := "SELECT code FROM confirmation_codes WHERE phone = $1 AND code = $2 AND created_at + INTERVAL '5 minutes' > NOW();"
	err := r.db.QueryRow(c, sql, dto.Phone, dto.Code).Scan(&code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return cerror.New(cerror.INVALID_CONFIRMATION_CODE, "invalid confirmation code")
		}
		return err
	}
	return nil
}
