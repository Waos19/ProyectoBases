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

// ==== helpers para NULL ====

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

// Podrías formatear fechas bonito, pero por ahora las mostramos crudas.
func t(nt sql.NullTime) string {
	if nt.Valid {
		return nt.Time.Format("2006-01-02")
	}
	return ""
}

// ==== 1. Listado de empleados básico ====

func (r *OracleRepo) Q1_ListadoEmpleados() error {
	rows, err := r.DB.Query(`
        SELECT 
            IDENTIFICACION,
            TRIM(
                NOMBRE_1 || ' ' ||
                COALESCE(NOMBRE_2, '') || ' ' ||
                APELLIDO_1 || ' ' ||
                COALESCE(APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            FECHA_N,
            LUGAR_N,
            SALARIO
        FROM EMPLEADOS
        ORDER BY IDENTIFICACION`)
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

func (r *OracleRepo) Q2_EmpleadosConCargo() error {
	rows, err := r.DB.Query(`
        SELECT 
            e.IDENTIFICACION,
            TRIM(
                e.NOMBRE_1 || ' ' ||
                COALESCE(e.NOMBRE_2, '') || ' ' ||
                e.APELLIDO_1 || ' ' ||
                COALESCE(e.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            h.CARGO,
            h.FECHA_INGRESO
        FROM EMPLEADOS e
        JOIN HISTORIAL_LABORAL h
          ON e.IDENTIFICACION = h.EMPLEADO
        ORDER BY e.IDENTIFICACION, h.FECHA_INGRESO DESC`)
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

func (r *OracleRepo) Q3_ListadoCargos() error {
	rows, err := r.DB.Query(`SELECT COD_CARGO, CARGO FROM CARGOS`)
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

func (r *OracleRepo) Q4_CargosAsignados() error {
	rows, err := r.DB.Query(`
        SELECT DISTINCT c.COD_CARGO, c.CARGO
        FROM CARGOS c
        JOIN HISTORIAL_LABORAL h
          ON h.CARGO = c.COD_CARGO`)
	// si quieres solo los actuales: AND h.ACTUAL = 'S'
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

func (r *OracleRepo) Q5_CargosAsignadosConCantidad() error {
	rows, err := r.DB.Query(`
        SELECT 
            c.CARGO,
            COUNT(*) AS CANTIDAD_ASIGNADO
        FROM CARGOS c
        JOIN HISTORIAL_LABORAL h
          ON h.CARGO = c.COD_CARGO
        GROUP BY c.CARGO`)
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

func (r *OracleRepo) Q6_HistorialLaboral() error {
	rows, err := r.DB.Query(`
        SELECT 
            e.IDENTIFICACION,
            TRIM(
                e.NOMBRE_1 || ' ' ||
                COALESCE(e.NOMBRE_2, '') || ' ' ||
                e.APELLIDO_1 || ' ' ||
                COALESCE(e.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            h.FECHA_INGRESO,
            h.FECHA_SALIDA,
            h.CARGO
        FROM EMPLEADOS e
        JOIN HISTORIAL_LABORAL h
          ON e.IDENTIFICACION = h.EMPLEADO
        ORDER BY e.IDENTIFICACION, h.FECHA_INGRESO`)
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

func (r *OracleRepo) Q7_CargosTecnico() error {
	rows, err := r.DB.Query(`
        SELECT COD_CARGO, CARGO 
        FROM CARGOS
        WHERE CARGO LIKE 'TECNICO%'`)
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

func (r *OracleRepo) Q8_CargosAuxiliar() error {
	rows, err := r.DB.Query(`
        SELECT COD_CARGO, CARGO 
        FROM CARGOS
        WHERE CARGO LIKE 'AUXILIAR%'`)
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

func (r *OracleRepo) Q9_CargosADSL() error {
	rows, err := r.DB.Query(`
        SELECT COD_CARGO, CARGO 
        FROM CARGOS
        WHERE CARGO LIKE '%ADSL%'`)
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

func (r *OracleRepo) Q10_EmpleadosActivos() error {
	rows, err := r.DB.Query(`
        SELECT 
            e.IDENTIFICACION,
            TRIM(
                e.NOMBRE_1 || ' ' ||
                COALESCE(e.NOMBRE_2, '') || ' ' ||
                e.APELLIDO_1 || ' ' ||
                COALESCE(e.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO
        FROM EMPLEADOS e
        JOIN HISTORIAL_LABORAL h
          ON e.IDENTIFICACION = h.EMPLEADO
        WHERE h.FECHA_SALIDA IS NULL`)
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

// ==== 11. Elementos asignados vigentes (ACTUAL='S') ====

func (r *OracleRepo) Q11_ElementosAsignadosVigentes() error {
	rows, err := r.DB.Query(`
        SELECT 
            TRIM(
                emp.NOMBRE_1 || ' ' ||
                COALESCE(emp.NOMBRE_2, '') || ' ' ||
                emp.APELLIDO_1 || ' ' ||
                COALESCE(emp.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            emp.IDENTIFICACION,
            ele.ELEMENTO
        FROM EMPLEADOS emp
        JOIN ELEMENTOS_ASIGNADOS asig
          ON emp.IDENTIFICACION = asig.EMPLEADO
        JOIN ELEMENTOS ele
          ON ele.CODIGO = asig.ELEMENTO
        WHERE asig.ACTUAL = 'S'
        ORDER BY emp.IDENTIFICACION ASC, ele.ELEMENTO ASC`)
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

func (r *OracleRepo) Q12_HistorialElementosAsignados() error {
	rows, err := r.DB.Query(`
        SELECT 
            TRIM(
                emp.NOMBRE_1 || ' ' ||
                COALESCE(emp.NOMBRE_2, '') || ' ' ||
                emp.APELLIDO_1 || ' ' ||
                COALESCE(emp.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            emp.IDENTIFICACION,
            ele.ELEMENTO,
            asig.CANTIDAD,
            asig.DURACION,
            asig.NUMERO AS TALLA
        FROM EMPLEADOS emp
        JOIN ELEMENTOS_ASIGNADOS asig
          ON emp.IDENTIFICACION = asig.EMPLEADO
        JOIN ELEMENTOS ele
          ON ele.CODIGO = asig.ELEMENTO
        ORDER BY emp.IDENTIFICACION ASC, ele.ELEMENTO ASC`)
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

func (r *OracleRepo) Q13_EntregasRealizadas() error {
	rows, err := r.DB.Query(`
        SELECT
            ent.ID_ENTREGA,
            TRIM(
                emp.NOMBRE_1 || ' ' ||
                COALESCE(emp.NOMBRE_2, '') || ' ' ||
                emp.APELLIDO_1 || ' ' ||
                COALESCE(emp.APELLIDO_2, '')
            ) AS NOMBRE_EMPLEADO,
            ele.ELEMENTO,
            ent.FECHA
        FROM ENTREGA_ELEMENTOS ent
        JOIN EMPLEADOS emp
          ON emp.IDENTIFICACION = ent.EMPLEADO
        JOIN ELEMENTOS ele
          ON ele.CODIGO = ent.ELEMENTO
        ORDER BY ent.FECHA DESC`)
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

func (r *OracleRepo) Q14_ElementosPorTrabajador() error {
	rows, err := r.DB.Query(`
        SELECT 
            ele.ELEMENTO,
            TRIM(
                emp.NOMBRE_1 || ' ' ||
                COALESCE(emp.NOMBRE_2, '') || ' ' ||
                emp.APELLIDO_1 || ' ' ||
                COALESCE(emp.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO
        FROM EMPLEADOS emp
        JOIN ELEMENTOS_ASIGNADOS asi
          ON emp.IDENTIFICACION = asi.EMPLEADO
        JOIN ELEMENTOS ele
          ON ele.CODIGO = asi.ELEMENTO`)
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

func (r *OracleRepo) Q15_TotalElementosPorTrabajador() error {
	rows, err := r.DB.Query(`
        SELECT
            TRIM(
                emp.NOMBRE_1 || ' ' ||
                COALESCE(emp.NOMBRE_2, '') || ' ' ||
                emp.APELLIDO_1 || ' ' ||
                COALESCE(emp.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            emp.IDENTIFICACION,
            ele.ELEMENTO,
            COUNT(ent.ID_ENTREGA) AS TOTAL_ENTREGADOS
        FROM ENTREGA_ELEMENTOS ent
        JOIN EMPLEADOS emp
          ON emp.IDENTIFICACION = ent.EMPLEADO
        JOIN ELEMENTOS ele
          ON ele.CODIGO = ent.ELEMENTO
        GROUP BY
            emp.NOMBRE_1, emp.NOMBRE_2, emp.APELLIDO_1, emp.APELLIDO_2,
            emp.IDENTIFICACION, ele.ELEMENTO
        ORDER BY emp.IDENTIFICACION ASC, ele.ELEMENTO ASC`)
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

func (r *OracleRepo) Q16_TotalElementos() error {
	rows, err := r.DB.Query(`
        SELECT
            ele.ELEMENTO,
            COUNT(ent.ID_ENTREGA) AS TOTAL_ENTREGADOS
        FROM ENTREGA_ELEMENTOS ent
        JOIN ELEMENTOS ele
          ON ele.CODIGO = ent.ELEMENTO
        GROUP BY ele.ELEMENTO
        ORDER BY ele.ELEMENTO ASC`)
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

func (r *OracleRepo) Q17_FichaTecnicaElementos() error {
	rows, err := r.DB.Query(`
        SELECT
            ELEMENTO,
            MATERIALES,
            MANTENIMIENTO,
            USOS,
            NORMA,
            ATENUACION,
            RUTA
        FROM ELEMENTOS
        ORDER BY ELEMENTO ASC`)
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

// ==== 18. Empleados y elementos pendientes por entregar (VIGENTE = 'N') ====

func (r *OracleRepo) Q18_ElementosPendientesEmpleado() error {
	rows, err := r.DB.Query(`
        SELECT
            TRIM(
                emp.NOMBRE_1 || ' ' ||
                COALESCE(emp.NOMBRE_2, '') || ' ' ||
                emp.APELLIDO_1 || ' ' ||
                COALESCE(emp.APELLIDO_2, '')
            ) AS NOMBRE_COMPLETO,
            emp.IDENTIFICACION,
            ele.ELEMENTO
        FROM ENTREGA_ELEMENTOS ent
        JOIN EMPLEADOS emp
          ON emp.IDENTIFICACION = ent.EMPLEADO
        JOIN ELEMENTOS ele
          ON ele.CODIGO = ent.ELEMENTO
        WHERE ent.VIGENTE = 'N'
        ORDER BY emp.IDENTIFICACION ASC, ele.ELEMENTO ASC`)
	if err != nil {
		return fmt.Errorf("q18: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n=== 18. Elementos pendientes por empleado (VIGENTE = 'N') ===")
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

func (r *OracleRepo) Q19_ElementoMasSolicitado() error {
	row := r.DB.QueryRow(`
        SELECT elemento, total_entregados
        FROM (
            SELECT 
                ele.ELEMENTO AS elemento,
                COUNT(ent.ID_ENTREGA) AS total_entregados
            FROM ENTREGA_ELEMENTOS ent
            JOIN ELEMENTOS ele
              ON ele.CODIGO = ent.ELEMENTO
            GROUP BY ele.ELEMENTO
            ORDER BY total_entregados DESC
        )
        WHERE ROWNUM = 1`)

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

func (r *OracleRepo) Q20_ElementosSegundoPeriodo2009() error {
	rows, err := r.DB.Query(`
        SELECT
            ele.ELEMENTO,
            COUNT(ent.ID_ENTREGA) AS TOTAL_ENTREGADOS
        FROM ENTREGA_ELEMENTOS ent
        JOIN ELEMENTOS ele
          ON ele.CODIGO = ent.ELEMENTO
        WHERE EXTRACT(YEAR FROM ent.FECHA) = 2009
          AND EXTRACT(MONTH FROM ent.FECHA) BETWEEN 7 AND 12
        GROUP BY ele.ELEMENTO
        ORDER BY ele.ELEMENTO ASC`)
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
