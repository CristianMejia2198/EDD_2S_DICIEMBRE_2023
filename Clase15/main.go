package main

import "Clase15/estructuras"

func main() {
	arbolito := &estructuras.ArbolMerkle{RaizMerkle: nil, BloqueDeDatos: nil, CantidadBloques: 0}
	arbolito.AgregarBloque("Aceptado", "Libro1", 201700918)
	arbolito.AgregarBloque("Rechazado", "Libro2", 201700918)
	arbolito.AgregarBloque("Aceptado", "Libro3", 201700918)
	arbolito.AgregarBloque("Aceptado", "Libro4", 201700918)
	arbolito.AgregarBloque("Rechazado", "Libro5", 201700918)
	arbolito.AgregarBloque("Aceptado", "Libro6", 201700918)
	arbolito.AgregarBloque("Aceptado", "Libro7", 201700918)
	arbolito.AgregarBloque("Aceptado", "Libro8", 201700918)

	arbolito.GenerarArbol()

	arbolito.Graficar()
}
