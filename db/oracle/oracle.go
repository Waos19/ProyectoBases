package oracle

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

type OracleRepo struct {
	DB *sql.DB
}

func OraConnect() (*OracleRepo, error) {
	dsn := "proyecto/sistemas@localhost:1521/XEDB"

	oracleDB, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión: %w", err)
	}

	if err := oracleDB.Ping(); err != nil {
		return nil, fmt.Errorf("error haciendo ping a Oracle: %w", err)
	}

	return &OracleRepo{DB: oracleDB}, nil
}

//
// =========================
//  FORMATO BONITO
// =========================
//

func printHeader(title string) {
	fmt.Println("\n========================================")
	fmt.Println(">>", title)
	fmt.Println("========================================")
}

func printLine() {
	fmt.Println("----------------------------------------")
}

//
// =========================
//  CONSULTAS
// =========================
//

func ListarServicios(db *sql.DB) ([]string, error) {
	printHeader("Lista de Servicios")

	rows, err := db.Query(`SELECT nombre FROM SERVICIO ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servicios []string
	for rows.Next() {
		var nombre string
		rows.Scan(&nombre)
		fmt.Println("Servicio:", nombre)
		servicios = append(servicios, nombre)
	}

	printLine()
	return servicios, nil
}

func ParticipantesActividadId(db *sql.DB, idActividad int) error {
	printHeader(fmt.Sprintf("Participantes de la Actividad %d", idActividad))

	rows, err := db.Query(`
        SELECT p.nombre, s.nombre AS servicio
        FROM PARTICIPACION pa
        JOIN PERSONAL p ON pa.id_personal = p.id_personal
        JOIN ACTIVIDAD a ON pa.id_actividad = a.id_actividad
        JOIN SERVICIO s ON a.id_servicio = s.id_servicio
        WHERE pa.id_actividad = :1
        ORDER BY p.nombre
    `, idActividad)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nombre string
		var servicio string
		rows.Scan(&nombre, &servicio)

		printLine()
		fmt.Println("Participante:", nombre)
		fmt.Println("Actividad:   ", servicio)
	}

	printLine()
	return nil
}

func ParticipantesActividad(db *sql.DB) error {
	printHeader("Todos los Participantes y sus Actividades")

	rows, err := db.Query(`
        SELECT p.nombre, s.nombre AS servicio
        FROM PARTICIPACION pa
        JOIN PERSONAL p ON pa.id_personal = p.id_personal
        JOIN ACTIVIDAD a ON pa.id_actividad = a.id_actividad
        JOIN SERVICIO s ON a.id_servicio = s.id_servicio
        ORDER BY s.nombre, p.nombre
    `)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nombre string
		var servicio string
		rows.Scan(&nombre, &servicio)

		printLine()
		fmt.Println("Participante:", nombre)
		fmt.Println("Actividad:   ", servicio)
	}

	printLine()
	return nil
}

func ActividadesCupoLleno(db *sql.DB) error {
	printHeader("Actividades con Cupo Lleno")

	rows, err := db.Query(`
        SELECT a.id_actividad, s.nombre, a.fecha
        FROM ACTIVIDAD a
        JOIN SERVICIO s ON a.id_servicio = s.id_servicio
        JOIN PARTICIPACION p ON a.id_actividad = p.id_actividad
        WHERE p.rol = 'participante'
        GROUP BY a.id_actividad, s.nombre, a.fecha, s.max_integrantes
        HAVING COUNT(p.id_personal) >= s.max_integrantes
    `)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var servicio string
		var fecha string
		rows.Scan(&id, &servicio, &fecha)

		printLine()
		fmt.Println("Actividad ID:", id)
		fmt.Println("Servicio:    ", servicio)
		fmt.Println("Fecha:       ", fecha)
	}

	printLine()
	return nil
}

func ElementosPorServicio(db *sql.DB) error {
	printHeader("Elementos Usados por Servicio")

	rows, err := db.Query(`
        SELECT s.nombre, e.nombre, SUM(ae.cantidad_usada)
        FROM ACTIVIDAD_ELEMENTO ae
        JOIN ACTIVIDAD a ON ae.id_actividad = a.id_actividad
        JOIN SERVICIO s ON a.id_servicio = s.id_servicio
        JOIN ELEMENTO e ON ae.id_elemento = e.id_elemento
        GROUP BY s.nombre, e.nombre
        ORDER BY s.nombre
    `)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var servicio, elemento string
		var total int
		rows.Scan(&servicio, &elemento, &total)

		printLine()
		fmt.Println("Servicio: ", servicio)
		fmt.Println("Elemento: ", elemento)
		fmt.Println("Cantidad: ", total)
	}

	printLine()
	return nil
}

func IngresosPorSemana(db *sql.DB) error {
	printHeader("Ingresos por Servicio y Semana")

	rows, err := db.Query(`
        SELECT s.nombre, TRUNC(a.fecha, 'IW'), SUM(pg.monto)
        FROM PAGO pg
        JOIN ACTIVIDAD a ON pg.id_actividad = a.id_actividad
        JOIN SERVICIO s ON a.id_servicio = s.id_servicio
        GROUP BY s.nombre, TRUNC(a.fecha, 'IW')
        ORDER BY s.nombre
    `)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var servicio, semana string
		var ingresos float64
		rows.Scan(&servicio, &semana, &ingresos)

		printLine()
		fmt.Println("Servicio: ", servicio)
		fmt.Println("Semana:   ", semana)
		fmt.Println("Ingresos: ", ingresos)
	}

	printLine()
	return nil
}

func SesionesYGastosGimnasio(db *sql.DB) error {
	printHeader("Sesiones y Gastos en Gimnasio")

	rows, err := db.Query(`
        SELECT per.nombre,
               COUNT(pa.id_actividad) AS sesiones,
               SUM(NVL(pg.monto,0)) AS total_pagado
        FROM PERSONAL per
        JOIN PARTICIPACION pa ON per.id_personal = pa.id_personal
        JOIN ACTIVIDAD a ON pa.id_actividad = a.id_actividad
        LEFT JOIN PAGO pg ON per.id_personal = pg.id_personal
        WHERE a.id_servicio = 1
          AND pa.rol = 'participante'
        GROUP BY per.nombre
        ORDER BY sesiones DESC
    `)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nombre string
		var sesiones int
		var pagado float64
		rows.Scan(&nombre, &sesiones, &pagado)

		printLine()
		fmt.Println("Nombre:   ", nombre)
		fmt.Println("Sesiones: ", sesiones)
		fmt.Println("Pagado:   ", pagado)
	}

	printLine()
	return nil
}
