package estructuras

type NodoListaAdyacencia struct {
	Siguiente *NodoListaAdyacencia
	Abajo     *NodoListaAdyacencia
	Valor     string
}

type PeticionGrafo struct {
	NombreArchivo string
}
