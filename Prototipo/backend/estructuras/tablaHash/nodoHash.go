package tablaHash

type Persona struct {
	Carnet   int
	Nombre   string
	Password string
	Cursos   []string
}

type NodoHash struct {
	Llave   int
	Persona *Persona
}
