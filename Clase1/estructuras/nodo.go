package estructuras

type Alumno struct {
	Carnet int
	Nombre string
	Curso  string
}

type NodoLista struct {
	Alumno    *Alumno
	Siguiente *NodoLista
}
