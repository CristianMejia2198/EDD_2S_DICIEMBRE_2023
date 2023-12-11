package main

import "Clase5/estructuras"

func main() {
	matrizPrueba := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Dato: "RAIZ"}}

	/*Primer Caso*/
	matrizPrueba.Insertar_Elemento(0, 0, "0,0")
	matrizPrueba.Insertar_Elemento(2, 2, "2,2")
	matrizPrueba.Insertar_Elemento(4, 4, "4,4")
	matrizPrueba.Insertar_Elemento(6, 6, "6,6")
	matrizPrueba.Insertar_Elemento(8, 8, "8,8")
	matrizPrueba.Reporte("MatrizCaso1.jpg")
	/*Segundo Caso*/
	matrizPrueba.Insertar_Elemento(0, 1, "0,1")
	matrizPrueba.Insertar_Elemento(2, 3, "2,3")
	matrizPrueba.Reporte("MatrizCaso2.jpg")
	/*Tercer Caso*/
	matrizPrueba.Insertar_Elemento(1, 2, "1,2")
	matrizPrueba.Insertar_Elemento(5, 0, "5,0")
	matrizPrueba.Reporte("MatrizCaso3.jpg")
	/*Cuarto Caso*/
	matrizPrueba.Insertar_Elemento(8, 4, "8,4")
	matrizPrueba.Insertar_Elemento(5, 2, "5,2")
	matrizPrueba.Reporte("MatrizFinal.jpg")
}
