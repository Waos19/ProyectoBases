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

func s(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func f(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0
}

func t(nt sql.NullTime) string {
	if nt.Valid {
		return nt.Time.Format("2006-01-02")
	}
	return ""
}

// ==== 1. Listado de empleados básico ====

func (r *PostgresRepo) Q1_ListadoEmpleados() error {
	rows, err := r.DB.Query(`
        SELECT 
            identificacion,
            TRIM(
                nombre_1 || ' ' ||
                COALESCE(nombre_2, '') || ' ' ||
                apellido_1 || ' ' ||
                COALESCE(apellido_2, '')
            ) AS nombre_completo,
            fecha_n,
            lugar_n,
            salario
        FROM empleados
        ORDER BY identificacion`)
	if err != nil {
		return fmt.Errorf("q1: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 1. Listado de empleados ===")

	for rows.Next() {
		var id int64
		var nombre string
		var fechaN sql.NullTime
		var lugar sql.NullString
		var salario sql.NullFloat64

		if err := rows.Scan(&id, &nombre, &fechaN, &lugar, &salario); err != nil {
			return err
		}

		fmt.Printf("ID: %d | Nombre: %s | FechaN: %s | LugarN: %s | Salario: %.2f\n",
			id, nombre, t(fechaN), s(lugar), f(salario))
	}
	return rows.Err()
}

// ==== 2. Empleados con cargo y fecha de ingreso ====

func (r *PostgresRepo) Q2_EmpleadosConCargo() error {
	rows, err := r.DB.Query(`
        SELECT 
            e.identificacion,
            TRIM(
                e.nombre_1 || ' ' ||
                COALESCE(e.nombre_2, '') || ' ' ||
                e.apellido_1 || ' ' ||
                COALESCE(e.apellido_2, '')
            ) AS nombre_completo,
            h.cargo,
            h.fecha_ingreso
        FROM empleados e
        JOIN historial_laboral h
          ON e.identificacion = h.empleado
        ORDER BY e.identificacion, h.fecha_ingreso DESC`)
	if err != nil {
		return fmt.Errorf("q2: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 2. Empleados con cargo y fecha de ingreso ===")

	for rows.Next() {
		var id, cargo int64
		var nombre string
		var fecha sql.NullTime

		if err := rows.Scan(&id, &nombre, &cargo, &fecha); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s | CargoID: %d | FechaIngreso: %s\n",
			id, nombre, cargo, t(fecha))
	}
	return rows.Err()
}

// ==== 3. Listado de cargos ====

func (r *PostgresRepo) Q3_ListadoCargos() error {
	rows, err := r.DB.Query(`SELECT cod_cargo, cargo FROM cargos`)
	if err != nil {
		return fmt.Errorf("q3: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 3. Listado de cargos ===")
	for rows.Next() {
		var id int64
		var cargo string
		if err := rows.Scan(&id, &cargo); err != nil {
			return err
		}
		fmt.Printf("CargoID: %d | Nombre: %s\n", id, cargo)
	}
	return rows.Err()
}

// ==== 4. Cargos actualmente asignados ====

func (r *PostgresRepo) Q4_CargosAsignados() error {
	rows, err := r.DB.Query(`
        SELECT DISTINCT c.cod_cargo, c.cargo
        FROM cargos c
        JOIN historial_laboral h
          ON h.cargo = c.cod_cargo`)
	// si quieres solo los actuales: AND h.actual = 'S'
	if err != nil {
		return fmt.Errorf("q4: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 4. Cargos asignados ===")
	for rows.Next() {
		var id int64
		var cargo string
		if err := rows.Scan(&id, &cargo); err != nil {
			return err
		}
		fmt.Printf("CargoID: %d | Cargo: %s\n", id, cargo)
	}
	return rows.Err()
}

// ==== 5. Cargos asignados y cantidad de empleados ====

func (r *PostgresRepo) Q5_CargosAsignadosConCantidad() error {
	rows, err := r.DB.Query(`
        SELECT 
            c.cargo,
            COUNT(*) AS cantidad_asignado
        FROM cargos c
        JOIN historial_laboral h
          ON h.cargo = c.cod_cargo
        GROUP BY c.cargo`)
	if err != nil {
		return fmt.Errorf("q5: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 5. Cargos y cantidad de empleados ===")
	for rows.Next() {
		var cargo string
		var cant int64
		if err := rows.Scan(&cargo, &cant); err != nil {
			return err
		}
		fmt.Printf("Cargo: %s | Cantidad empleados: %d\n", cargo, cant)
	}
	return rows.Err()
}

// ==== 6. Historial laboral por empleado ====

func (r *PostgresRepo) Q6_HistorialLaboral() error {
	rows, err := r.DB.Query(`
        SELECT 
            e.identificacion,
            TRIM(
                e.nombre_1 || ' ' ||
                COALESCE(e.nombre_2, '') || ' ' ||
                e.apellido_1 || ' ' ||
                COALESCE(e.apellido_2, '')
            ) AS nombre_completo,
            h.fecha_ingreso,
            h.fecha_salida,
            h.cargo
        FROM empleados e
        JOIN historial_laboral h
          ON e.identificacion = h.empleado
        ORDER BY e.identificacion, h.fecha_ingreso`)
	if err != nil {
		return fmt.Errorf("q6: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 6. Historial laboral ===")
	for rows.Next() {
		var id, cargo int64
		var nombre string
		var fIng, fSal sql.NullTime

		if err := rows.Scan(&id, &nombre, &fIng, &fSal, &cargo); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s | CargoID: %d | Ingreso: %s | Salida: %s\n",
			id, nombre, cargo, t(fIng), t(fSal))
	}
	return rows.Err()
}

// ==== 7. Cargos que comienzan por TECNICO ====

func (r *PostgresRepo) Q7_CargosTecnico() error {
	rows, err := r.DB.Query(`
        SELECT cod_cargo, cargo 
        FROM cargos
        WHERE cargo LIKE 'tecnico%'`)
	if err != nil {
		return fmt.Errorf("q7: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 7. Cargos que empiezan por 'TECNICO' ===")
	for rows.Next() {
		var id int64
		var cargo string
		if err := rows.Scan(&id, &cargo); err != nil {
			return err
		}
		fmt.Printf("CargoID: %d | Cargo: %s\n", id, cargo)
	}
	return rows.Err()
}

// ==== 8. Cargos que comienzan por AUXILIAR ====

func (r *PostgresRepo) Q8_CargosAuxiliar() error {
	rows, err := r.DB.Query(`
        SELECT cod_cargo, cargo 
        FROM cargos
        WHERE cargo LIKE 'auxiliar%'`)
	if err != nil {
		return fmt.Errorf("q8: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 8. Cargos que empiezan por 'AUXILIAR' ===")
	for rows.Next() {
		var id int64
		var cargo string
		if err := rows.Scan(&id, &cargo); err != nil {
			return err
		}
		fmt.Printf("CargoID: %d | Cargo: %s\n", id, cargo)
	}
	return rows.Err()
}

// ==== 9. Cargos que contienen 'ADSL' ====

func (r *PostgresRepo) Q9_CargosADSL() error {
	rows, err := r.DB.Query(`
        SELECT cod_cargo, cargo 
        FROM cargos
        WHERE cargo LIKE '%adsl%'`)
	if err != nil {
		return fmt.Errorf("q9: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 9. Cargos que contienen 'ADSL' ===")
	for rows.Next() {
		var id int64
		var cargo string
		if err := rows.Scan(&id, &cargo); err != nil {
			return err
		}
		fmt.Printf("CargoID: %d | Cargo: %s\n", id, cargo)
	}
	return rows.Err()
}

// ==== 10. Empleados cuya fecha de salida es nula ====

func (r *PostgresRepo) Q10_EmpleadosActivos() error {
	rows, err := r.DB.Query(`
        SELECT 
            e.identificacion,
            TRIM(
                e.nombre_1 || ' ' ||
                COALESCE(e.nombre_2, '') || ' ' ||
                e.apellido_1 || ' ' ||
                COALESCE(e.apellido_2, '')
            ) AS nombre_completo
        FROM empleados e
        JOIN historial_laboral h
          ON e.identificacion = h.empleado
        WHERE h.fecha_salida IS NULL`)
	if err != nil {
		return fmt.Errorf("q10: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 10. Empleados sin fecha de salida (activos) ===")
	for rows.Next() {
		var id int64
		var nombre string
		if err := rows.Scan(&id, &nombre); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s\n", id, nombre)
	}
	return rows.Err()
}

// ==== 11. Elementos asignados vigentes (actual='S') ====

func (r *PostgresRepo) Q11_ElementosAsignadosVigentes() error {
	rows, err := r.DB.Query(`
        SELECT 
            TRIM(
                emp.nombre_1 || ' ' ||
                COALESCE(emp.nombre_2, '') || ' ' ||
                emp.apellido_1 || ' ' ||
                COALESCE(emp.apellido_2, '')
            ) AS nombre_completo,
            emp.identificacion,
            ele.elemento
        FROM empleados emp
        JOIN elementos_asignados asig
          ON emp.identificacion = asig.empleado
        JOIN elementos ele
          ON ele.codigo = asig.elemento
        WHERE asig.actual = 's'
        ORDER BY emp.identificacion ASC, ele.elemento ASC`)
	if err != nil {
		return fmt.Errorf("q11: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 11. Elementos asignados vigentes ===")
	for rows.Next() {
		var nombre string
		var id int64
		var elem string
		if err := rows.Scan(&nombre, &id, &elem); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s | Elemento: %s\n", id, nombre, elem)
	}
	return rows.Err()
}

// ==== 12. Historial de elementos asignados ====

func (r *PostgresRepo) Q12_HistorialElementosAsignados() error {
	rows, err := r.DB.Query(`
        SELECT 
            TRIM(
                emp.nombre_1 || ' ' ||
                COALESCE(emp.nombre_2, '') || ' ' ||
                emp.apellido_1 || ' ' ||
                COALESCE(emp.apellido_2, '')
            ) AS nombre_completo,
            emp.identificacion,
            ele.elemento,
            asig.cantidad,
            asig.duracion,
            asig.numero AS talla
        FROM empleados emp
        JOIN elementos_asignados asig
          ON emp.identificacion = asig.empleado
        JOIN elementos ele
          ON ele.codigo = asig.elemento
        ORDER BY emp.identificacion ASC, ele.elemento ASC`)
	if err != nil {
		return fmt.Errorf("q12: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 12. Historial de elementos asignados ===")
	for rows.Next() {
		var nombre, elem, talla string
		var id int64
		var cant, dur sql.NullInt64

		if err := rows.Scan(&nombre, &id, &elem, &cant, &dur, &talla); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s | Elemento: %s | Cantidad: %v | Duración: %v | Talla: %s\n",
			id, nombre, elem, cant, dur, talla)
	}
	return rows.Err()
}

// ==== 13. Entregas realizadas ====

func (r *PostgresRepo) Q13_EntregasRealizadas() error {
	rows, err := r.DB.Query(`
        SELECT
            ent.id_entrega,
            TRIM(
                emp.nombre_1 || ' ' ||
                COALESCE(emp.nombre_2, '') || ' ' ||
                emp.apellido_1 || ' ' ||
                COALESCE(emp.apellido_2, '')
            ) AS nombre_empleado,
            ele.elemento,
            ent.fecha
        FROM entrega_elementos ent
        JOIN empleados emp
          ON emp.identificacion = ent.empleado
        JOIN elementos ele
          ON ele.codigo = ent.elemento
        ORDER BY ent.fecha DESC`)
	if err != nil {
		return fmt.Errorf("q13: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 13. Entregas realizadas ===")
	for rows.Next() {
		var idEntrega int64
		var nombre, elem string
		var fecha sql.NullTime

		if err := rows.Scan(&idEntrega, &nombre, &elem, &fecha); err != nil {
			return err
		}
		fmt.Printf("EntregaID: %d | Empleado: %s | Elemento: %s | Fecha: %s\n",
			idEntrega, nombre, elem, t(fecha))
	}
	return rows.Err()
}

// ==== 14. Elementos entregados a cada trabajador ====

func (r *PostgresRepo) Q14_ElementosPorTrabajador() error {
	rows, err := r.DB.Query(`
        SELECT 
            ele.elemento,
            TRIM(
                emp.nombre_1 || ' ' ||
                COALESCE(emp.nombre_2, '') || ' ' ||
                emp.apellido_1 || ' ' ||
                COALESCE(emp.apellido_2, '')
            ) AS nombre_completo
        FROM empleados emp
        JOIN elementos_asignados asi
          ON emp.identificacion = asi.empleado
        JOIN elementos ele
          ON ele.codigo = asi.elemento`)
	if err != nil {
		return fmt.Errorf("q14: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 14. Elementos entregados a cada trabajador ===")
	for rows.Next() {
		var elem, nombre string
		if err := rows.Scan(&elem, &nombre); err != nil {
			return err
		}
		fmt.Printf("Empleado: %s | Elemento: %s\n", nombre, elem)
	}
	return rows.Err()
}

// ==== 15. Total de elementos entregados a cada trabajador, agrupados por elemento ====

func (r *PostgresRepo) Q15_TotalElementosPorTrabajador() error {
	rows, err := r.DB.Query(`
        SELECT
            TRIM(
                emp.nombre_1 || ' ' ||
                COALESCE(emp.nombre_2, '') || ' ' ||
                emp.apellido_1 || ' ' ||
                COALESCE(emp.apellido_2, '')
            ) AS nombre_completo,
            emp.identificacion,
            ele.elemento,
            COUNT(ent.id_entrega) AS total_entregados
        FROM entrega_elementos ent
        JOIN empleados emp
          ON emp.identificacion = ent.empleado
        JOIN elementos ele
          ON ele.codigo = ent.elemento
        GROUP BY emp.nombre_1, emp.nombre_2, emp.apellido_1, emp.apellido_2,
                 emp.identificacion, ele.elemento
        ORDER BY emp.identificacion ASC, ele.elemento ASC`)
	if err != nil {
		return fmt.Errorf("q15: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 15. Total elementos entregados por trabajador y elemento ===")
	for rows.Next() {
		var nombre, elem string
		var id, total int64
		if err := rows.Scan(&nombre, &id, &elem, &total); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s | Elemento: %s | Total: %d\n",
			id, nombre, elem, total)
	}
	return rows.Err()
}

// ==== 16. Total de elementos entregados, agrupados por elemento ====

func (r *PostgresRepo) Q16_TotalElementos() error {
	rows, err := r.DB.Query(`
        SELECT
            ele.elemento,
            COUNT(ent.id_entrega) AS total_entregados
        FROM entrega_elementos ent
        JOIN elementos ele
          ON ele.codigo = ent.elemento
        GROUP BY ele.elemento
        ORDER BY ele.elemento ASC`)
	if err != nil {
		return fmt.Errorf("q16: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 16. Total elementos entregados por elemento ===")
	for rows.Next() {
		var elem string
		var total int64
		if err := rows.Scan(&elem, &total); err != nil {
			return err
		}
		fmt.Printf("Elemento: %s | Total entregados: %d\n", elem, total)
	}
	return rows.Err()
}

// ==== 17. Ficha técnica de cada elemento ====

func (r *PostgresRepo) Q17_FichaTecnicaElementos() error {
	rows, err := r.DB.Query(`
        SELECT
            elemento,
            materiales,
            mantenimiento,
            usos,
            norma,
            atenuacion,
            ruta
        FROM elementos
        ORDER BY elemento ASC`)
	if err != nil {
		return fmt.Errorf("q17: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 17. Ficha técnica de elementos ===")
	for rows.Next() {
		var elem, materiales, mant, usos, norma, atenuacion, ruta sql.NullString
		if err := rows.Scan(&elem, &materiales, &mant, &usos, &norma, &atenuacion, &ruta); err != nil {
			return err
		}
		fmt.Printf("Elemento: %s\n  Materiales: %s\n  Mantenimiento: %s\n  Usos: %s\n  Norma: %s\n  Atenuación: %s\n  Ruta: %s\n\n",
			s(elem), s(materiales), s(mant), s(usos), s(norma), s(atenuacion), s(ruta))
	}
	return rows.Err()
}

// ==== 18. Empleados y elementos pendientes por entregar (vigente = 'N') ====

func (r *PostgresRepo) Q18_ElementosPendientesEmpleado() error {
	rows, err := r.DB.Query(`
        SELECT
            TRIM(
                emp.nombre_1 || ' ' ||
                COALESCE(emp.nombre_2, '') || ' ' ||
                emp.apellido_1 || ' ' ||
                COALESCE(emp.apellido_2, '')
            ) AS nombre_completo,
            emp.identificacion,
            ele.elemento
        FROM entrega_elementos ent
        JOIN empleados emp
          ON emp.identificacion = ent.empleado
        JOIN elementos ele
          ON ele.codigo = ent.elemento
        WHERE ent.vigente = 'n'
        ORDER BY emp.identificacion ASC, ele.elemento ASC`)
	if err != nil {
		return fmt.Errorf("q18: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 18. Elementos pendientes por empleado (vigente = 'N') ===")
	for rows.Next() {
		var nombre, elem string
		var id int64
		if err := rows.Scan(&nombre, &id, &elem); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Nombre: %s | Elemento pendiente: %s\n",
			id, nombre, elem)
	}
	return rows.Err()
}

// ==== 19. Elemento más solicitado ====

func (r *PostgresRepo) Q19_ElementoMasSolicitado() error {
	row := r.DB.QueryRow(`
        SELECT elemento, total_entregados
        FROM (
            SELECT 
                ele.elemento AS elemento,
                COUNT(ent.id_entrega) AS total_entregados
            FROM entrega_elementos ent
            JOIN elementos ele
              ON ele.codigo = ent.elemento
            GROUP BY ele.elemento
            ORDER BY total_entregados DESC
            LIMIT 1
        ) sub`)

	fmt.Println("\n=== 19. Elemento más solicitado ===")

	var elem string
	var total int64
	if err := row.Scan(&elem, &total); err != nil {
		return fmt.Errorf("q19: %w", err)
	}
	fmt.Printf("Elemento: %s | Total entregados: %d\n", elem, total)
	return nil
}

// ==== 20. Elementos entregados en segundo periodo 2009 ====

func (r *PostgresRepo) Q20_ElementosSegundoPeriodo2009() error {
	rows, err := r.DB.Query(`
        SELECT
            ele.elemento,
            COUNT(ent.id_entrega) AS total_entregados
        FROM entrega_elementos ent
        JOIN elementos ele
          ON ele.codigo = ent.elemento
        WHERE EXTRACT(YEAR FROM ent.fecha) = 2009
          AND EXTRACT(MONTH FROM ent.fecha) BETWEEN 7 AND 12
        GROUP BY ele.elemento
        ORDER BY ele.elemento ASC`)
	if err != nil {
		return fmt.Errorf("q20: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 20. Elementos entregados (2º periodo 2009) ===")
	for rows.Next() {
		var elem string
		var total int64
		if err := rows.Scan(&elem, &total); err != nil {
			return err
		}
		fmt.Printf("Elemento: %s | Total entregados: %d\n", elem, total)
	}
	return rows.Err()
}
