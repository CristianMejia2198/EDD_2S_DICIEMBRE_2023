package estructuras

type Tutor struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoCola struct {
	Tutor     *Tutor
	Prioridad int
	Siguiente *NodoCola
}
