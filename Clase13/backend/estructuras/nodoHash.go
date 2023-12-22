package estructuras

type Persona struct {
	Carnet   int
	Nombre   string
	Password string
}

type NodoHash struct {
	Llave   int
	Persona *Persona
}
