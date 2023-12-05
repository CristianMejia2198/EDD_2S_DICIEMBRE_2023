package main

import "Clase1/estructuras"

func main() {
	lista := estructuras.ListaSimple{Inicio: nil, Final: nil, Longitud: 0}
	// ListaSimple lista = new ListaSimple()
	lista.Insertar(202400000, "Estudiante 1", "0770")
	lista.Insertar(202400001, "Estudiante 2", "0770")
	lista.Insertar(202400002, "Estudiante 3", "0770")
	lista.Insertar(202400003, "Estudiante 4", "0770")
	lista.Insertar(202400004, "Estudiante 5", "0770")
	lista.Insertar(202400005, "Estudiante 6", "0770")
	lista.Mostrar()
}
