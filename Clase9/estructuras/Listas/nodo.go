package Listas

type Alumno struct {
	Carnet int
	Nombre string
}

type Tutores struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoListaDoble struct {
	Alumno    *Alumno
	Siguiente *NodoListaDoble
	Anterior  *NodoListaDoble
}

type NodoListaCircular struct {
	Tutor     *Tutores
	Siguiente *NodoListaCircular
	Anterior  *NodoListaCircular
}
