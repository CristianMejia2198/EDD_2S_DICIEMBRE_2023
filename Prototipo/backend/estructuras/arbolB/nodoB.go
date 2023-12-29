package arbolB

type Libro struct {
	Nombre    string
	Contenido string
	Estado    int
}
type Publicacion struct {
	Contenido string
}

type Tutores struct {
	Carnet        int
	Nombre        string
	Curso         string
	Password      string
	Libros        []*Libro
	Publicaciones []*Publicacion
}

type NodoB struct {
	Valor     *Tutores
	Siguiente *NodoB
	Anterior  *NodoB
	Izquierdo *RamaB
	Derecho   *RamaB
}
