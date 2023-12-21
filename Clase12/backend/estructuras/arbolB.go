package estructuras

import (
	"fmt"
	"strconv"
)

type ArbolB struct {
	Raiz  *RamaB
	Orden int
}

func (a *ArbolB) insertar_rama(nodo *NodoB, rama *RamaB) *NodoB { // 20,
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

func (a *ArbolB) dividir(rama *RamaB) *NodoB { //rama = ( 10 | 15 | 20 )
	val := &NodoB{Valor: -9999}                                 // Nodo Temporal
	aux := rama.Primero                                         //Auxiliar para recorrer la lista
	rderecha := &RamaB{Primero: nil, Contador: 0, Hoja: true}   //ramas temporales
	rizquierda := &RamaB{Primero: nil, Contador: 0, Hoja: true} //ramas temporales
	contador := 0                                               // Ayuda para saber en que nodo estoy actualmente
	for aux != nil {                                            //Recorrer una rama
		contador++        //1
		if contador < 2 { // 15 | 20 | 25 -> i15d
			temp := &NodoB{Valor: aux.Valor} // NodoB = (valor = 10)
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

func (a *ArbolB) Insertar(valor int) { //15
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

/***************************************/
func (a *ArbolB) Graficar(nombre string) {
	cadena := ""
	nombre_archivo := "./" + nombre + ".dot"
	nombre_imagen := nombre + ".jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol { \nnode[shape=record]\n"
		cadena += a.grafo(a.Raiz.Primero)
		cadena += a.conexionRamas(a.Raiz.Primero)
		cadena += "}"
	}
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolB) grafo(rama *NodoB) string {
	dot := ""
	if rama != nil {
		dot += a.grafoRamas(rama)
		aux := rama
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += a.grafo(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += a.grafo(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

func (a *ArbolB) grafoRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		dot = dot + "R" + strconv.Itoa(rama.Valor) + "[label=\""
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot = dot + "<C" + strconv.Itoa(r) + ">|"
				r++
			}
			if aux.Siguiente != nil {
				dot = dot + strconv.Itoa(aux.Valor) + "|"
			} else {
				dot = dot + strconv.Itoa(aux.Valor)
				if aux.Derecho != nil {
					dot = dot + "|<C" + strconv.Itoa(r) + ">"
				}
			}
			aux = aux.Siguiente
		}
		dot = dot + "\"];\n"
	}
	return dot
}

func (a *ArbolB) conexionRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		actual := "R" + strconv.Itoa(rama.Valor)
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + strconv.Itoa(aux.Izquierdo.Primero.Valor) + ";\n"
				r++
				dot += a.conexionRamas(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + strconv.Itoa(aux.Derecho.Primero.Valor) + ";\n"
					r++
					dot += a.conexionRamas(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

/********************/
func (a *ArbolB) Buscar(numero int) {
	buscarElemento := a.buscarArbol(a.Raiz.Primero, numero)
	if buscarElemento != nil {
		fmt.Println("Se encontro el elemento", buscarElemento)
	} else {
		fmt.Println("No se encontro")
	}
}

func (a *ArbolB) buscarArbol(raiz *NodoB, numero int) *NodoB {
	var valorEncontrado *NodoB
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				valorEncontrado = a.buscarArbol(aux.Izquierdo.Primero, numero)
			}
			if aux.Valor == numero {
				fmt.Println("Se enconto el valor: ", numero)
				valorEncontrado = aux
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					valorEncontrado = a.buscarArbol(aux.Derecho.Primero, numero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return valorEncontrado
}
