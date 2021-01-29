package psql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Repository interface {
	GetConnection() *pgx.Conn
}

type repository struct {
	Client *pgx.Conn
}

func New(ctx context.Context, dsn string) (Repository, error) {
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	repo := repository{Client: db}
	if err = repo.Client.Ping(ctx); err != nil {
		return &repo, err
	}

	return &repo, nil
}

func (c repository) GetConnection() *pgx.Conn {
	return c.Client
}
