package main

import (
	"Clase3/estructuras"
	"fmt"
)

var ColaPrioridad *estructuras.Cola = &estructuras.Cola{Inicio: nil, Longitud: 0}

func main() {
	//ColaPrioridad := estructuras.Cola{Inicio: nil, Longitud: 0}

	opcion := 0
	salir := false

	for !salir {
		MenuAdmin()
		fmt.Scanln(&opcion)
		if opcion == 1 {
			//llamar a su funcion de leer csv
			ColaPrioridad.EncolarPrioridad(201700001, "Estudiante 1", "0770", 95)
			ColaPrioridad.EncolarPrioridad(201700002, "Estudiante 2", "0771", 64)
			ColaPrioridad.EncolarPrioridad(201700003, "Estudiante 3", "0772", 75)
			ColaPrioridad.EncolarPrioridad(201700004, "Estudiante 4", "0773", 80)
			ColaPrioridad.EncolarPrioridad(201700005, "Estudiante 5", "0774", 100)
			ColaPrioridad.EncolarPrioridad(201700006, "Estudiante 6", "0775", 50)
			ColaPrioridad.EncolarPrioridad(201700007, "Estudiante 7", "0781", 71)
			ColaPrioridad.EncolarPrioridad(201700008, "Estudiante 8", "0771", 63)
			fmt.Println("Se cargo a la cola de prioridad")
		} else if opcion == 2 {
			MenuPrincipal()
		} else {
			salir = true
		}
	}
}

func MenuPrincipal() {
	opcion := 0
	salir := false

	for !salir {
		ColaPrioridad.Primero()
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			fmt.Println("Se agrego a la lista al estudiante ", ColaPrioridad.Inicio.Tutor.Carnet)
			ColaPrioridad.Descolar()
		} else if opcion == 2 {
			ColaPrioridad.Descolar()
		} else {
			salir = true
		}
	}

}

func MenuAdmin() {

	fmt.Println("1. Cargar Tutores")
	fmt.Println("2. Control de Estudiantes Tutores")
	fmt.Println("3. Cerrar Programa")
}
