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

	// --- DEMO Oracle: insertar, actualizar, borrar y listar ---

	// if err := repoOra.InsertarEmpleadoDemo(); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := repoOra.ActualizarEmpleadoDemo(); err != nil {
	// 	log.Fatal(err)
	// }

	// // Por si quieres ver al empleado DEMO en el listado antes de borrarlo:
	// if err := repoOra.Q1_ListadoEmpleados(); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := repoOra.BorrarEmpleadoDemo(); err != nil {
	// 	log.Fatal(err)
	// }

	// // Listado final para mostrar que ya no está el DEMO
	// if err := repoOra.Q1_ListadoEmpleados(); err != nil {
	// 	log.Fatal(err)
	// }

	// repoOra.Q1_ListadoEmpleados()
	// repoOra.Q2_EmpleadosConCargo()
	// repoOra.Q3_ListadoCargos()
	// repoOra.Q4_CargosAsignados()
	// repoOra.Q5_CargosAsignadosConCantidad()
	// repoOra.Q6_HistorialLaboral()
	// repoOra.Q7_CargosTecnico()
	// repoOra.Q8_CargosAuxiliar()
	// repoOra.Q9_CargosADSL()
	// repoOra.Q10_EmpleadosActivos()
	// repoOra.Q11_ElementosAsignadosVigentes()
	// repoOra.Q12_HistorialElementosAsignados()
	// repoOra.Q13_EntregasRealizadas()
	// repoOra.Q14_ElementosPorTrabajador()
	// repoOra.Q15_TotalElementosPorTrabajador()
	// repoOra.Q16_TotalElementos()
	// repoOra.Q17_FichaTecnicaElementos()
	// repoOra.Q18_ElementosPendientesEmpleado()
	// repoOra.Q19_ElementoMasSolicitado()
	// repoOra.Q20_ElementosSegundoPeriodo2009()

	repoPos.Q1_ListadoEmpleados()
	repoPos.Q2_EmpleadosConCargo()
	repoPos.Q3_ListadoCargos()
	repoPos.Q4_CargosAsignados()
	repoPos.Q5_CargosAsignadosConCantidad()
	repoPos.Q6_HistorialLaboral()
	repoPos.Q7_CargosTecnico()
	repoPos.Q8_CargosAuxiliar()
	repoPos.Q9_CargosADSL()
	repoPos.Q10_EmpleadosActivos()
	repoPos.Q11_ElementosAsignadosVigentes()
	repoPos.Q12_HistorialElementosAsignados()
	repoPos.Q13_EntregasRealizadas()
	repoPos.Q14_ElementosPorTrabajador()
	repoPos.Q15_TotalElementosPorTrabajador()
	repoPos.Q16_TotalElementos()
	repoPos.Q17_FichaTecnicaElementos()
	repoPos.Q18_ElementosPendientesEmpleado()
	repoPos.Q19_ElementoMasSolicitado()
	repoPos.Q20_ElementosSegundoPeriodo2009()

}
