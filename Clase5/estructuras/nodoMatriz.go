package estructuras

type NodoMatriz struct {
	Siguiente *NodoMatriz //Derecha
	Anterior  *NodoMatriz //Izquierda
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Dato      string
}

/*
[01 02 03 04]
[05 06 07 08]
[09 10 11 12]
[13 14 15 16]

*/
