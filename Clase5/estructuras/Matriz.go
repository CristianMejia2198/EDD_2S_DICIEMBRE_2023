package estructuras

import (
	"fmt"
	"strconv"
)

type Matriz struct {
	Raiz *NodoMatriz
}

func (m *Matriz) buscarColumna(x int) *NodoMatriz {
	aux := m.Raiz
	for aux != nil {
		if aux.PosX == x {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

func (m *Matriz) buscarFila(y int) *NodoMatriz {
	aux := m.Raiz
	for aux != nil {
		if aux.PosY == y {
			return aux
		}
		aux = aux.Abajo
	}
	return nil
}

/* y = 3 -> no existe cabecera o columna (debo crearla)
Raiz := &NodoMatriz{PosX: -1, PosY: -1, Dato: "RAIZ"}
Raiz	C1 C2 C3 C5
	F1  P
	F2
	F3
	F5
*/

func (m *Matriz) insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for {
		if temp.PosX == nuevoNodo.PosX {
			temp.PosY = nuevoNodo.PosY
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosX > nuevoNodo.PosX {
			piv = true
			break
		}
		if temp.Siguiente != nil {
			temp = temp.Siguiente
		} else {
			break
		}
	}
	if piv {
		nuevoNodo.Siguiente = temp
		nuevoNodo.Anterior = temp.Anterior
		temp.Anterior.Siguiente = nuevoNodo
		temp.Anterior = nuevoNodo
	} else {
		nuevoNodo.Anterior = temp
		temp.Siguiente = nuevoNodo
	}
	return nuevoNodo
}

func (m *Matriz) nuevaColumna(x int) *NodoMatriz { //nuevaColumna(x int, carnet_Tutor int)
	col := "C" + strconv.Itoa(x)
	nuevoNodo := &NodoMatriz{PosX: x, PosY: -1, Dato: col}
	columna := m.insertarColumna(nuevoNodo, m.Raiz)
	return columna
}

func (m *Matriz) insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for {
		if temp.PosY == nuevoNodo.PosY {
			temp.PosX = nuevoNodo.PosX
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosY > nuevoNodo.PosY {
			piv = true
			break
		}
		if temp.Abajo != nil {
			temp = temp.Abajo
		} else {
			break
		}
	}
	if piv {
		/*
			nuevoNodo.Siguiente = temp
			nuevoNodo.Anterior = temp.Anterior
			temp.Anterior.Siguiente = nuevoNodo
			temp.Anterior = nuevoNodo
		*/
		nuevoNodo.Abajo = temp
		nuevoNodo.Arriba = temp.Arriba
		temp.Arriba.Abajo = nuevoNodo
		temp.Arriba = nuevoNodo
	} else {
		nuevoNodo.Arriba = temp
		temp.Abajo = nuevoNodo
	}
	return nuevoNodo
}

func (m *Matriz) nuevaFila(y int) *NodoMatriz { //nuevaFila(y int, carnet_estudiante int)
	fil := "F" + strconv.Itoa(y) // strconv.Itoa(carnet)
	nuevoNodo := &NodoMatriz{PosX: -1, PosY: y, Dato: fil}
	fila := m.insertarFila(nuevoNodo, m.Raiz)
	return fila
}

func (m *Matriz) Insertar_Elemento(x int, y int, valor string) {
	nuevoNodo := &NodoMatriz{PosX: x, PosY: y, Dato: valor}
	nodoColumna := m.buscarColumna(x)
	nodoFila := m.buscarFila(y)
	/*
		1. Columna y Fila no existen
		2. Columna si existe pero Fila no
		3. Columna no existe pero Fila si
		4. Columna y Fila si existen
	*/
	/*
		Raiz	C1 (1, -1)
			F1  P
		  (-1,1)
	*/

	/*
		Raiz	C1 C2
			F1  P
			F2  Q  X
	*/

	if nodoColumna == nil && nodoFila == nil { // (1,1)
		nodoColumna = m.nuevaColumna(x)
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila == nil { //(1, 2)
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna == nil && nodoFila != nil { //(2, 2)
		nodoColumna = m.nuevaColumna(x)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila != nil { //(1,2)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else {
		fmt.Println("ERROR!!!")
	}
}

func (m *Matriz) Reporte(nombre string) {
	texto := ""
	nombre_archivo := "./matriz.dot"
	nombre_imagen := nombre
	aux1 := m.Raiz
	aux2 := m.Raiz
	aux3 := m.Raiz
	if aux1 != nil {
		texto = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n"
		/** Creacion de los nodos actuales */
		for aux1 != nil {
			texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Dato + "\" ,rankdir=LR,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
			aux1 = aux1.Siguiente
		}
		texto += "}"
		for aux2 != nil {
			aux1 = aux2
			texto += "{rank=same; \n"
			for aux1 != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Dato + "\" ,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
				aux1 = aux1.Siguiente
			}
			texto += "}"
			aux2 = aux2.Abajo
		}
		/** Conexiones entre los nodos de la matriz */
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Siguiente != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Siguiente.PosX+1) + strconv.Itoa(aux1.Siguiente.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Abajo != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Abajo.PosX+1) + strconv.Itoa(aux1.Abajo.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Abajo
			}
			aux2 = aux2.Siguiente
		}
		texto += "}"
	} else {
		texto = "No hay elementos en la matriz"
	}
	//fmt.Println(texto)
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
