package main

import (
	"log"
	"proybases/db/oracle"
	"proybases/db/postgresdb"
)

func main() {
	// Conexión a Oracle
	repoOra, err := oracle.OraConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer repoOra.DB.Close()

	// Conexión a PostgreSQL
	repoPos, err := postgresdb.PostConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer repoPos.DB.Close()
}
