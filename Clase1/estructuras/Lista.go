package estructuras

import "fmt"

type ListaSimple struct {
	Inicio   *NodoLista
	Final    *NodoLista
	Longitud int
}

// clase alumno, cambiarNombre, this.Longitud, self.Longitud
func (l *ListaSimple) Insertar(carnet int, nombre string, curso string) {
	alumno := &Alumno{Carnet: carnet, Nombre: nombre, Curso: curso}
	if l.Longitud == 0 {
		nuevo := &NodoLista{Alumno: alumno, Siguiente: nil}
		l.Inicio = nuevo // 0X123456AB
		l.Final = nuevo  // 0X123456AB
		l.Longitud++
	} else {
		//Bucle -> while
		/*
			aux := l.Inicio
			for aux.Siguiente != nil {
				aux = aux.Siguiente
			}
			aux.siguiente = &NodoLista{Alumno: alumno, Siguiente: nil}
			l.Longitud++
		*/
		nuevo := &NodoLista{Alumno: alumno, Siguiente: nil}
		l.Final.Siguiente = nuevo
		l.Final = nuevo
		l.Longitud++
	}
}

func (l *ListaSimple) Mostrar() {
	aux := l.Inicio
	for aux != nil {
		fmt.Println(aux.Alumno.Carnet, "->", aux.Alumno.Nombre, "->", aux.Alumno.Curso, " | ")
		aux = aux.Siguiente
	}
}

// Clase alumno, hablar, participar,
