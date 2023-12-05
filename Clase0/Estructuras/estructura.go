package Estructuras

type Alumno struct {
	Carnet int
	Nombre string
}

type Nodo struct {
	alumno    *Alumno
	siguiente *Nodo
}

type Lista struct {
	Inicio *Nodo
}

//punteros
/*

 * declarar variables de tipo puntero -> referencia a un espacio, 0x1535DAC
 & referencia al espacio de memoria


*/
