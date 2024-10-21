package repository

import "context"

func (r *repo) CodesCountLastHourByIP(c context.Context, ip string) (int, error) {
	var count int
	sql := "SELECT COUNT(ip) as count FROM confirmation_codes WHERE ip = $1 AND created_at >= (NOW() AT TIME ZONE 'UTC') - INTERVAL '5 minutes';"
	err := r.db.QueryRow(c, sql, ip).Scan(&count)
	return count, err
}
