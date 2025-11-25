package postgresdb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresRepo struct {
	DB *sql.DB
}

func PostConnect() (*PostgresRepo, error) {
	dsn := "postgres://postgres@localhost:5432/postgres"

	PostDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión: %w", err)
	}

	// Probar conexión
	if err := PostDB.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("error haciendo ping a PostgreSQL: %w", err)
	}

	return &PostgresRepo{DB: PostDB}, nil
}
