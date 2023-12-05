package main

import (
	"Clase0/Estructuras"
	"fmt"
)

func main() {
	// Declaracion basica de Go -> var nombrevariable tipo = valor
	var mivariable int = 0
	fmt.Println(mivariable)

	//Declaracion corta de Go -> nombredevariable := valor
	mivariable2 := 2.3
	fmt.Println(mivariable2)

	//Manejo de variables o atributos publicos o privados
	//(publico primer caracter mayuscula)
	//(privado primer caracter minuscula)
	alumno1 := Estructuras.Alumno{Carnet: 0, Nombre: "alumno1"}
	fmt.Println(alumno1)

	//Memoria estatica, creando los espacios de memoria
	milista := [4]int{1, 2}
	fmt.Println(milista)

	//Memoria dinamica
	lista := &Estructuras.Lista{Inicio: nil}
	fmt.Println(&lista)

	var x int = 5
	y := &x
	fmt.Println(x)
	(*y)++ // +1 0xc00001 -> obtengo el valor x = 5
	fmt.Println(x)
	z := x
	z++ // 0xc00002 -> obtengo el valor de z = x
	fmt.Println(x)
	fmt.Println(z)
	// mivariable = 5
	fmt.Println("Hola")
}
