package repository

import (
	"authorization_service/internal/core/configuration"
	"authorization_service/internal/infrastructure"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func New(config *configuration.DBConfig) (infrastructure.Repository, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s", config.User, config.Pass, config.Host, config.Name)
	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &repo{db: conn}, nil
}

func (r *repo) Close() {
	r.db.Close()
}
