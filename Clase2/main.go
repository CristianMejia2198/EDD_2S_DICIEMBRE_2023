package main

import "Clase2/estructuras"

func main() {
	lista := estructuras.ListaDoble{Inicio: nil, Longitud: 0}
	lista.LeerArchivo("Estudiante.csv")
	lista.Mostrar()
}
