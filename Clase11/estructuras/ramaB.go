package estructuras

type RamaB struct {
	Primero  *NodoB
	Hoja     bool
	Contador int
}

func (r *RamaB) Insertar(nuevoNodo *NodoB) {
	if r.Primero == nil { // 10
		r.Primero = nuevoNodo
		r.Contador++
	} else {
		if nuevoNodo.Valor < r.Primero.Valor { // 7 | 10
			nuevoNodo.Siguiente = r.Primero
			nuevoNodo.Derecho = r.Primero.Izquierdo
			r.Primero.Anterior = nuevoNodo
			r.Primero = nuevoNodo
			r.Contador++
		} else if r.Primero.Siguiente != nil { // 7 | 10 -> 9
			if r.Primero.Siguiente.Valor > nuevoNodo.Valor {
				nuevoNodo.Siguiente = r.Primero.Siguiente
				nuevoNodo.Anterior = r.Primero
				nuevoNodo.Derecho = r.Primero.Siguiente.Izquierdo
				nuevoNodo.Izquierdo = r.Primero.Derecho
				r.Primero.Siguiente.Anterior = nuevoNodo
				r.Primero.Siguiente = nuevoNodo
				r.Contador++
			} else { // 7 | 10 -> 15
				aux := r.Primero.Siguiente
				nuevoNodo.Anterior = aux
				nuevoNodo.Izquierdo = aux.Derecho
				aux.Siguiente = nuevoNodo
				r.Contador++
			}
		}
	}
}
