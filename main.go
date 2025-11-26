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

	oracle.ListarServicios(repoOra.DB)

	oracle.ParticipantesActividad(repoOra.DB)

	oracle.ActividadesCupoLleno(repoOra.DB)

	oracle.ElementosPorServicio(repoOra.DB)

	oracle.IngresosPorSemana(repoOra.DB)

	oracle.SesionesYGastosGimnasio(repoOra.DB)
}
