package repository

import "context"

func (r *repo) DeleteCodesByPhone(c context.Context, phone string) error {
	sql := "DELETE FROM confirmation_codes WHERE phone = $1"
	return r.db.QueryRow(c, sql, phone).Scan()
}
