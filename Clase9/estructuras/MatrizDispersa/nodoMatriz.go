package MatrizDispersa

type Dato struct {
	Carnet_Tutor      int
	Carnet_Estudiante int
	Curso             string
}

/*
Cabeceras
	Carnet_Tutor -> dato -> cualquier carnet
	Cartnet_Alumno -> 0
Filas
	Carnet_Alumno -> dato -> cualquier carnet
	Carnet_Tutor -> 0
Nodos Internos
	Carnet_Tutor -> dato -> cualquier carnet
	Carnet_Alumno -> dato -> cualquier carnet
	Curso
*/

type NodoMatriz struct {
	Siguiente *NodoMatriz //Derecha
	Anterior  *NodoMatriz //Izquierda
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Dato      *Dato
}
