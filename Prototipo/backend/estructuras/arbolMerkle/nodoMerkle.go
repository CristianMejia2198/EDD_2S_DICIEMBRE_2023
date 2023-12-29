package arbolMerkle

type NodoMerkle struct {
	Izquierda *NodoMerkle
	Derecha   *NodoMerkle
	Bloque    *NodoBloqueDatos
	Valor     string
}

type InformacionBloque struct {
	Fecha  string
	Accion string
	Nombre string
	Tutor  int
}

type NodoBloqueDatos struct {
	Siguiente *NodoBloqueDatos
	Anterior  *NodoBloqueDatos
	Valor     *InformacionBloque
}
