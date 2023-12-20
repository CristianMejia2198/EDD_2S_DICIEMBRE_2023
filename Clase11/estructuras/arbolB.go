package estructuras

type ArbolB struct {
	Raiz  *RamaB
	Orden int
}

func (a *ArbolB) insertar_rama(nodo *NodoB, rama *RamaB) *NodoB { // 20 | 15
	if rama.Hoja {
		rama.Insertar(nodo)
		if rama.Contador == a.Orden {
			return a.dividir(rama)
		} else {
			return nil
		}
	} else {
		temp := rama.Primero
		for temp != nil {
			if nodo.Valor == temp.Valor { //comparar si son igual y no agregarlo
				return nil
			} else if nodo.Valor < temp.Valor {
				obj := a.insertar_rama(nodo, temp.Izquierdo)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			} else if temp.Siguiente == nil {
				obj := a.insertar_rama(nodo, temp.Derecho)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			}
			temp = temp.Siguiente
		}
	}
	return nil
}

func (a *ArbolB) dividir(rama *RamaB) *NodoB { //rama = ( 15 | 20 | 25 )
	val := &NodoB{Valor: -9999}                                 // Nodo Temporal
	aux := rama.Primero                                         //Auxiliar para recorrer la lista
	rderecha := &RamaB{Primero: nil, Contador: 0, Hoja: true}   //ramas temporales
	rizquierda := &RamaB{Primero: nil, Contador: 0, Hoja: true} //ramas temporales
	contador := 0                                               // Ayuda para saber en que nodo estoy actualmente
	for aux != nil {                                            //Recorrer una rama
		contador++        //1
		if contador < 2 { // 15 | 20 | 25 -> i15d
			temp := &NodoB{Valor: aux.Valor} // NodoB = (valor = 15)
			temp.Izquierdo = aux.Izquierdo   // temp.izquierdo = 15.izquierdo
			if contador == 1 {
				temp.Derecho = aux.Siguiente.Izquierdo
			}
			if temp.Derecho != nil && temp.Izquierdo != nil { //Comparo, si son diferentes de null -> dividi un nodo raiz
				rizquierda.Hoja = false //Noo es nodo hoja, sino raiz
			}
			rizquierda.Insertar(temp)
		} else if contador == 2 {
			val.Valor = aux.Valor
		} else {
			temp := &NodoB{Valor: aux.Valor}
			temp.Izquierdo = aux.Izquierdo
			temp.Derecho = aux.Derecho
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rderecha.Hoja = false
			}
			rderecha.Insertar(temp)
		}
		aux = aux.Siguiente
	}
	nuevo := &NodoB{Valor: val.Valor}
	nuevo.Derecho = rderecha
	nuevo.Izquierdo = rizquierda
	/*
			rama = ( 15 | 20 | 25 )
			20
		15		25
	*/
	return nuevo
}

func (a *ArbolB) Insertar(valor int) {
	nuevoNodo := &NodoB{Valor: valor}
	if a.Raiz == nil {
		a.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
		a.Raiz.Insertar(nuevoNodo)
	} else {
		obj := a.insertar_rama(nuevoNodo, a.Raiz)
		if obj != nil {
			a.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
			a.Raiz.Insertar(obj)
			a.Raiz.Hoja = false
		}
	}
}
