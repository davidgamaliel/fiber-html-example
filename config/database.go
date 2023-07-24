package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDBPool(user, passwd, host, port, schema string, logger Logger) (*pgxpool.Pool, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?&pool_min_conns=1&pool_max_conn_lifetime=10s&pool_max_conn_idle_time=5s",
		user,
		passwd,
		host,
		port,
		schema,
	)

	db, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		logger.Error("failed to connect DB")
		return nil, err
	}

	return db, nil
}
