package ColaPrioridad

type Tutores struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoCola struct {
	Tutor     *Tutores
	Siguiente *NodoCola
	Prioridad int
}
