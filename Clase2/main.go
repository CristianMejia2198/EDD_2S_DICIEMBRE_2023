package main

import (
	"Clase2/estructuras"
	"fmt"
)

func factorialIterativo(n int) int {
	resultado := 1
	for i := 1; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

func factorialRecursivo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorialRecursivo(n-1)
}

func main() {
	lista := estructuras.ListaDoble{Inicio: nil, Longitud: 0}
	lista.LeerArchivo("Estudiante.csv")
	lista.Mostrar()
	fmt.Println(factorialIterativo(5))
	fmt.Println(factorialRecursivo(5))
}
