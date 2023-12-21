package main

import "Clase11/estructuras"

func main() {
	arbolitoB := &estructuras.ArbolB{Raiz: nil, Orden: 3}
	arbolitoB.Insertar(6)
	arbolitoB.Insertar(11)
	arbolitoB.Insertar(5)
	arbolitoB.Insertar(4)
	arbolitoB.Insertar(8)
	arbolitoB.Insertar(9)
	arbolitoB.Insertar(12)
	arbolitoB.Insertar(21)
	arbolitoB.Insertar(14)
	arbolitoB.Insertar(10)
	arbolitoB.Insertar(19)
	arbolitoB.Insertar(28)
	arbolitoB.Insertar(3)
	arbolitoB.Insertar(17)
	arbolitoB.Insertar(32)
	arbolitoB.Insertar(15)
	arbolitoB.Insertar(16)
	arbolitoB.Insertar(26)
	arbolitoB.Insertar(27)
	arbolitoB.Graficar()
	arbolitoB.Buscar(17)
}
