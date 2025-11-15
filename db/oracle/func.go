package oracle

import (
	"fmt"
)

// ==== DEMO: insertar, actualizar y borrar un empleado ====

func (r *OracleRepo) InsertarEmpleadoDemo() error {
	fmt.Println("\n>>> Insertando empleado DEMO...")

	_, err := r.DB.Exec(`
        INSERT INTO EMPLEADOS (
            IDENTIFICACION,
            TIPO,
            NOMBRE_1,
            NOMBRE_2,
            APELLIDO_1,
            APELLIDO_2,
            SEXO,
            FECHA_N,
            LUGAR_N,
            DIRECCION,
            TELEFONO,
            EMAIL,
            SALARIO,
            ACTIVO
        )
        VALUES (
            99999999,              -- IDENTIFICACION
            'CC',                  -- TIPO
            'DEMO',                -- NOMBRE_1
            NULL,                  -- NOMBRE_2
            'PRUEBA',              -- APELLIDO_1
            NULL,                  -- APELLIDO_2
            'M',                   -- SEXO
            DATE '1990-01-01',     -- FECHA_N
            'CIUDAD DEMO',         -- LUGAR_N
            'CALLE FALSA 123',     -- DIRECCION
            '3000000000',          -- TELEFONO
            'demo@correo.com',     -- EMAIL
            1000000,               -- SALARIO
            'S'                    -- ACTIVO
        )
    `)
	if err != nil {
		return fmt.Errorf("insertar empleado demo: %w", err)
	}

	fmt.Println("Empleado DEMO insertado (ID 99999999).")
	return nil
}

func (r *OracleRepo) ActualizarEmpleadoDemo() error {
	fmt.Println("\n>>> Actualizando empleado DEMO...")

	_, err := r.DB.Exec(`
        UPDATE EMPLEADOS
        SET SALARIO = SALARIO + 50000
        WHERE IDENTIFICACION = 99999999
    `)
	if err != nil {
		return fmt.Errorf("actualizar empleado demo: %w", err)
	}

	fmt.Println("Empleado DEMO actualizado (salario +50000).")
	return nil
}

func (r *OracleRepo) BorrarEmpleadoDemo() error {
	fmt.Println("\n>>> Borrando empleado DEMO...")

	_, err := r.DB.Exec(`
        DELETE FROM EMPLEADOS
        WHERE IDENTIFICACION = 99999999
    `)
	if err != nil {
		return fmt.Errorf("borrar empleado demo: %w", err)
	}

	fmt.Println("Empleado DEMO borrado.")
	return nil
}
