package arbolMerkle

import (
	"backend/estructuras/Peticiones"
	"encoding/hex"
	"math"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"
)

type ArbolMerkle struct {
	RaizMerkle      *NodoMerkle
	BloqueDeDatos   *NodoBloqueDatos
	CantidadBloques int
}

func fechaActual() string {
	now := time.Now()
	formato := "02-01-2006::15:04:05"
	fechahoraFormato := now.Format(formato) // 27-12-2023::12:02:40
	return fechahoraFormato
}

func (a *ArbolMerkle) AgregarBloque(estado string, nombreLibro string, carnet int) {
	nuevoRegistro := &InformacionBloque{Fecha: fechaActual(), Accion: estado, Nombre: nombreLibro, Tutor: carnet}
	nuevoBloque := &NodoBloqueDatos{Valor: nuevoRegistro}
	if a.BloqueDeDatos == nil {
		a.BloqueDeDatos = nuevoBloque
		a.CantidadBloques++
	} else {
		aux := a.BloqueDeDatos
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoBloque.Anterior = aux
		aux.Siguiente = nuevoBloque
		a.CantidadBloques++
	}
}

func (a *ArbolMerkle) GenerarArbol() {
	nivel := 1
	for int(math.Pow(2, float64(nivel))) < a.CantidadBloques {
		nivel++
	}
	for i := a.CantidadBloques; i < int(math.Pow(2, float64(nivel))); i++ {
		a.AgregarBloque(strconv.Itoa(i), "nulo", 0)
	}
	/*
		♫ -> ☼ -> ☼ -> ☼ -> ☼ -> nulo -> nulo -> nulo
	*/
	a.generarHash()
}

func (a *ArbolMerkle) generarHash() {
	var arrayNodos []*NodoMerkle
	aux := a.BloqueDeDatos
	for aux != nil {
		contanetacion := aux.Valor.Fecha + aux.Valor.Accion + aux.Valor.Nombre + strconv.Itoa(aux.Valor.Tutor)
		encriptado := a.encriptarSha3(contanetacion)
		nodoHoja := &NodoMerkle{Valor: encriptado, Bloque: aux}
		arrayNodos = append(arrayNodos, nodoHoja)
		aux = aux.Siguiente
	}
	a.RaizMerkle = a.crearArbol(arrayNodos)
}

func (a *ArbolMerkle) crearArbol(arrayNodos []*NodoMerkle) *NodoMerkle {
	var auxNodos []*NodoMerkle
	var raiz *NodoMerkle
	if len(arrayNodos) == 2 {
		encriptado := a.encriptarSha3(arrayNodos[0].Valor + arrayNodos[1].Valor)
		raiz = &NodoMerkle{Valor: encriptado}
		raiz.Izquierda = arrayNodos[0]
		raiz.Derecha = arrayNodos[1]
		return raiz
	} else {
		for i := 0; i < len(arrayNodos); i += 2 {
			encriptado := a.encriptarSha3(arrayNodos[i].Valor + arrayNodos[i+1].Valor)
			nodoRaiz := &NodoMerkle{Valor: encriptado}
			nodoRaiz.Izquierda = arrayNodos[i]
			nodoRaiz.Derecha = arrayNodos[i+1]
			auxNodos = append(auxNodos, nodoRaiz)
		}
		return a.crearArbol(auxNodos)
	}
}

func (a *ArbolMerkle) encriptarSha3(cadena string) string {
	hash := sha3.New256()
	hash.Write([]byte(cadena))
	encriptacion := hex.EncodeToString(hash.Sum(nil))
	return encriptacion
}

/*******************************************/
func (a *ArbolMerkle) Graficar() {
	cadena := ""
	nombre_archivo := "./Reporte/arbolMerkle.dot"
	nombre_imagen := "./Reporte/arbolMerkle.jpg"
	if a.RaizMerkle != nil {
		cadena += "digraph arbol { node [shape=box];"
		cadena += a.retornarValoresArbol(a.RaizMerkle, 0)
		cadena += "}"
	}
	Peticiones.CrearArchivo(nombre_archivo)
	Peticiones.EscribirArchivo(cadena, nombre_archivo)
	Peticiones.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolMerkle) retornarValoresArbol(raiz *NodoMerkle, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" [dir=back];\n"
		if raiz.Izquierda != nil && raiz.Derecha != nil {
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierda.Valor[:20]) + "\"" + " -> " + "\"" + (raiz.Derecha.Valor[:20]) + "\"" + " [style=invis]}; \n"
		}
	}
	if raiz.Bloque != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" -> "
		cadena += "\""
		cadena += raiz.Bloque.Valor.Fecha + "\n" + raiz.Bloque.Valor.Accion + "\n" + raiz.Bloque.Valor.Nombre + "\n" + strconv.Itoa(raiz.Bloque.Valor.Tutor)
		cadena += "\" [dir=back];\n "
	}
	return cadena
}
