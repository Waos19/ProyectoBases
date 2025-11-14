package oracle

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

type OracleRepo struct {
	DB *sql.DB
}

func DbConnect() (*OracleRepo, error) {
	//Usuario
	dsn := "estudiante/sistemas@localhost:1521/XEDB"

	oracleDB, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión: %w", err)
	}

	if err := oracleDB.Ping(); err != nil {
		return nil, fmt.Errorf("error haciendo ping a Oracle: %w", err)
	}

	return &OracleRepo{DB: oracleDB}, nil
}
