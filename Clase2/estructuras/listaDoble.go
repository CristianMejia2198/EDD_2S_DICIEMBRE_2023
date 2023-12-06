package estructuras

import "fmt"

type ListaDoble struct {
	Inicio   *NodoListaDoble
	Longitud int
}

func (l *ListaDoble) Agregar(carnet int, nombre string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodoListaDoble{Alumno: nuevoAlumno, Siguiente: nil, Anterior: nil}

	if l.Longitud == 0 {
		l.Inicio = nuevoNodo
		l.Longitud++
	} else { // 1 -> 2 -> 3 -> null
		aux := l.Inicio            // 1 -> 0x123
		for aux.Siguiente != nil { //aux = 3
			aux = aux.Siguiente
		} // aux = 3
		nuevoNodo.Anterior = aux
		aux.Siguiente = nuevoNodo
		l.Longitud++
	}
}

//Funcion de Agregar para una Lista Circular Doblemente Enlazada
func (l *ListaDoble) AgregarOrdenado(carnet int, nombre string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodoListaDoble{Alumno: nuevoAlumno, Siguiente: nil, Anterior: nil}

	if l.Longitud == 0 {
		l.Inicio = nuevoNodo
		l.Inicio.Anterior = nuevoNodo
		l.Inicio.Siguiente = nuevoNodo
		l.Longitud++
	} else {
		aux := l.Inicio
		contador := 1
		for contador < l.Longitud {
			if l.Inicio.Alumno.Carnet > carnet {
				nuevoNodo.Siguiente = l.Inicio
				nuevoNodo.Anterior = l.Inicio.Anterior
				l.Inicio.Anterior = nuevoNodo
				l.Inicio = nuevoNodo
				l.Longitud++
				return
			}
			if aux.Alumno.Carnet < carnet {
				aux = aux.Siguiente
			} else {
				nuevoNodo.Anterior = aux.Anterior
				aux.Anterior.Siguiente = nuevoNodo
				nuevoNodo.Siguiente = aux
				aux.Anterior = nuevoNodo
				l.Longitud++
				return
			}
			contador++
		}
		if aux.Alumno.Carnet > carnet {
			nuevoNodo.Siguiente = aux
			nuevoNodo.Anterior = aux.Anterior
			aux.Anterior.Siguiente = nuevoNodo
			aux.Anterior = nuevoNodo
			l.Longitud++
			return
		}
		nuevoNodo.Anterior = aux
		nuevoNodo.Siguiente = l.Inicio
		aux.Siguiente = nuevoNodo
		l.Longitud++
	}
}

func (l *ListaDoble) Mostrar() {
	aux := l.Inicio
	for aux.Siguiente != nil {
		fmt.Println(aux.Alumno.Carnet, " -> ", aux.Alumno.Nombre)
		aux = aux.Siguiente
	}
	fmt.Println(aux.Alumno.Carnet, " -> ", aux.Alumno.Nombre)
	fmt.Println("Descendente")
	for aux != nil {
		fmt.Println(aux.Alumno.Carnet, " -> ", aux.Alumno.Nombre)
		aux = aux.Anterior
	}
}

func (l *ListaDoble) MostrarV1() {
	aux := l.Inicio
	contador := 1
	for contador <= l.Longitud {
		fmt.Println(aux.Alumno.Carnet, " -> ", aux.Alumno.Nombre)
		aux = aux.Siguiente
		contador++
	}

}
