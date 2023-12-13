package main

import (
	"Clase7/estructuras"
	"fmt"
)

func main() {
	arbolBinario := estructuras.ArbolABB{Raiz: nil}
	arbolBinario.InsertarElemento("0773")
	arbolBinario.InsertarElemento("0771")
	arbolBinario.InsertarElemento("0770")
	arbolBinario.InsertarElemento("0772")
	arbolBinario.InsertarElemento("0775")
	arbolBinario.InsertarElemento("0774")
	arbolBinario.InsertarElemento("0777")
	arbolBinario.Recorridos()
	arbolBinario.Graficar()
	buscarElemento := arbolBinario.Busqueda("0781")
	if buscarElemento {
		fmt.Println("Existe el curso")
	} else {
		fmt.Println("No existe")
	}
}
