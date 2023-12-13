package estructuras

import (
	"fmt"
	"strconv"
)

type ArbolABB struct {
	Raiz *NodoArbol
}

func (a *ArbolABB) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor > nuevoNodo.Valor {
			raiz.Izquierdo = a.insertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.insertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	return raiz
}

func (a *ArbolABB) InsertarElemento(valor string) {
	nuevoNodo := &NodoArbol{Valor: valor}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *ArbolABB) recorridoInorden(raiz *NodoArbol) {
	//izquierdo -> raiz -> derecho
	if raiz != nil {
		if raiz.Izquierdo != nil {
			a.recorridoInorden(raiz.Izquierdo)
			fmt.Print("->")
		}
		fmt.Print(raiz.Valor, " ")
		if raiz.Derecho != nil {
			fmt.Print("->")
			a.recorridoInorden(raiz.Derecho)
		}
	}
}

func (a *ArbolABB) recorridoPreorden(raiz *NodoArbol) {
	//Raiz -> izquierdo -> derecho
	if raiz != nil {
		fmt.Print(raiz.Valor, " ")
		if raiz.Izquierdo != nil {
			fmt.Print("->")
			a.recorridoPreorden(raiz.Izquierdo)
		}
		if raiz.Derecho != nil {
			fmt.Print("->")
			a.recorridoPreorden(raiz.Derecho)
		}
	}
}

func (a *ArbolABB) recorridoPostorden(raiz *NodoArbol) {
	//izquierdo -> derecho -> raiz
	if raiz != nil {
		if raiz.Izquierdo != nil {
			a.recorridoPostorden(raiz.Izquierdo)
			fmt.Print("->")
		}
		if raiz.Derecho != nil {
			a.recorridoPostorden(raiz.Derecho)
			fmt.Print("->")
		}
		fmt.Print(raiz.Valor, " ")
	}
}

func (a *ArbolABB) Recorridos() {
	a.recorridoInorden(a.Raiz)
	fmt.Println("")
	a.recorridoPostorden(a.Raiz)
	fmt.Println("")
	a.recorridoPreorden(a.Raiz)
}

/*
			*
		+		15
	12		10

	+ 12 10
*/

func (a *ArbolABB) busqueda_arbol(valor string, raiz *NodoArbol) *NodoArbol {
	var valorEncontro *NodoArbol
	if raiz != nil {
		if raiz.Valor == valor {
			valorEncontro = raiz
		} else {
			if raiz.Valor > valor {
				valorEncontro = a.busqueda_arbol(valor, raiz.Izquierdo)
			} else {
				valorEncontro = a.busqueda_arbol(valor, raiz.Derecho)
			}
		}
	}
	return valorEncontro
}

func (a *ArbolABB) Busqueda(valor string) bool {
	buscarElemento := a.busqueda_arbol(valor, a.Raiz)
	if buscarElemento != nil {
		return true
	}
	return false
}

// Reporte Grafico
func (a *ArbolABB) Graficar() {
	cadena := ""
	nombre_archivo := "./arbolABB.dot"
	nombre_imagen := "arbolABB.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol{ "
		cadena += a.retornarValoresArbol(a.Raiz, 0)
		cadena += "}"
	}
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolABB) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor
		cadena += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "\"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}
