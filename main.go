package main

import (
	"context"
	"fmt"
	"proybases/db/oracle"
	"proybases/db/postgresdb"
)

func main() {
	repo, err := oracle.DbConnect()
	if err != nil {
		panic(err)
	}
	defer repo.DB.Close()

	var version string
	repo.DB.QueryRowContext(context.Background(),
		"SELECT banner FROM v$version WHERE ROWNUM = 1").Scan(&version)

	fmt.Println("Oracle version:", version)

	repoPos, err := postgresdb.PostConnect()
	if err != nil {
		panic(err)
	}
	defer repo.DB.Close()

	var version2 string
	repoPos.DB.QueryRowContext(context.Background(),
		"SELECT version()").Scan(&version2)

	fmt.Println("PostgreSQL version:", version2)
}
