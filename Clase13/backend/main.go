package main

import (
	"Clase13/estructuras"
	"fmt"
)

func main() {
	tablitaHash := estructuras.TablaHash{Tabla: make(map[int]estructuras.NodoHash), Capacidad: 7, Utilizacion: 0}
	/*tablitaHash.Insertar(201700918, "Cristian", "123")
	tablitaHash.Insertar(201712345, "Cristian", "123")
	tablitaHash.Insertar(201789369, "Cristian", "123")
	tablitaHash.Insertar(201700236, "Cristian", "123")
	tablitaHash.Insertar(201759369, "Cristian", "123")*/
	tablitaHash.LeerCSV("EstudiantesF1.csv")
	fmt.Println("Capacidad: ", tablitaHash.Capacidad)
	for i := 0; i < tablitaHash.Capacidad; i++ {
		if usuario, existe := tablitaHash.Tabla[i]; existe {
			fmt.Println("Posicion: ", i, " Carnet: ", usuario.Persona.Carnet)
		}
	} // 0,1,1,2,3,5,8,13,21,34

	fmt.Println("---------------------------------------------------")
	for i := 0; i < tablitaHash.Capacidad; i++ {
		if usuario, existe := tablitaHash.Tabla[i]; existe {
			fmt.Println("Llave: ", usuario.Llave, " Carnet: ", usuario.Persona.Carnet)
		}
	}
}
