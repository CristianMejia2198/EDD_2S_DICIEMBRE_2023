package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
Tabla = {
	3: nodoHash3,
	4: nodoHash4
}
*/

type TablaHash struct {
	Tabla       map[int]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) calculoIndice(carnet int) int {
	var numeros []int //201700001 -> [2 0 1 7 0 0 0 0 1]
	// Convertir el numero en array
	for {
		if carnet > 0 {
			digito := carnet % 10
			numeros = append([]int{digito}, numeros...)
			carnet = carnet / 10
		} else {
			break
		}
	}

	//Convertir array de numeros en Codigo ascii -> [1]
	var numeros_ascii []rune
	for _, numero := range numeros {
		valor := rune(numero + 48)
		numeros_ascii = append(numeros_ascii, valor)
	}

	//Sumar cada numero ascii
	final := 0
	for _, numero_ascii := range numeros_ascii {
		final += int(numero_ascii)
	}

	indice := final % t.Capacidad
	return indice
}

func (t *TablaHash) capacidadTabla() {
	auxCap := float64(t.Capacidad) * 0.6 // 7 * 0.6 --- 4
	if t.Utilizacion > int(auxCap) {
		auxAnterior := t.Capacidad
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar(auxAnterior)
	}

}

func (t *TablaHash) nuevaCapacidad() int {
	contador := 0
	a, b := 0, 1
	for contador < 50 {
		contador++
		if a > t.Capacidad {
			return a
		}
		a, b = b, a+b
	}
	return a
}

func (t *TablaHash) reInsertar(capacidadAnterior int) {
	auxTabla := t.Tabla
	t.Tabla = make(map[int]NodoHash)
	for i := 0; i < capacidadAnterior; i++ {
		if usuario, existe := auxTabla[i]; existe {
			t.Insertar(usuario.Persona.Carnet, usuario.Persona.Nombre, usuario.Persona.Password)
		}
	}
}

func (t *TablaHash) reCalculoIndice(carnet int, contador int) int {
	nuevoIndice := t.calculoIndice(carnet) + (contador * contador) // 2 + 36 = 38 -> capacidad actual 7
	return t.nuevoIndice(nuevoIndice)
}

func (t *TablaHash) nuevoIndice(nuevoIndice int) int {
	nuevoPosicion := 0
	if nuevoIndice < t.Capacidad {
		nuevoPosicion = nuevoIndice
	} else {
		nuevoPosicion = nuevoIndice - t.Capacidad
		nuevoPosicion = t.nuevoIndice(nuevoPosicion)
	}
	return nuevoPosicion
}

func (t *TablaHash) Insertar(carnet int, nombre string, password string) {
	indice := t.calculoIndice(carnet)
	nuevoNodo := &NodoHash{Llave: indice, Persona: &Persona{Carnet: carnet, Nombre: nombre, Password: password}}
	if indice < t.Capacidad {
		// 5
		if _, existe := t.Tabla[indice]; !existe {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidadTabla()
		} else {
			contador := 1
			indice = t.reCalculoIndice(carnet, contador)
			for {
				if _, existe := t.Tabla[indice]; existe {
					contador++
					indice = t.reCalculoIndice(carnet, contador)
				} else {
					nuevoNodo.Llave = indice
					t.Tabla[indice] = *nuevoNodo
					t.Utilizacion++
					t.capacidadTabla()
					break
				}
			}
		}
	}
}

func (t *TablaHash) LeerCSV(ruta string) {
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
		t.Insertar(valor, linea[1], "1234")
	}
}
