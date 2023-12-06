package estructuras

type Alumno struct {
	Carnet int
	Nombre string
}

type NodoListaDoble struct {
	Alumno    *Alumno
	Siguiente *NodoListaDoble
	Anterior  *NodoListaDoble
}
