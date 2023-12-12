package Listas

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

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
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoNodo.Anterior = aux
		aux.Siguiente = nuevoNodo
		l.Longitud++
	}
}

func (l *ListaDoble) Buscar(carnet string, password string) bool {
	if l.Longitud == 0 {
		return false
	} else {
		aux := l.Inicio
		for aux != nil {
			if strconv.Itoa(aux.Alumno.Carnet) == carnet && carnet == password {
				return true
			}
			aux = aux.Siguiente
		}
	}

	return false
}

func (l *ListaDoble) LeerCSV(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		l.Agregar(valor, linea[1])
	}
}
