package main

import (
	"Clase4/estructuras/ColaPrioridad"
	"Clase4/estructuras/Listas"
	"fmt"
)

/*
		import Lista

import lista	importa cola
*/

var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var colaPrioridad *ColaPrioridad.Cola = &ColaPrioridad.Cola{Primero: nil, Longitud: 0}

/*
row-major -> recorrer por filas
column-major -> recorrer por columnas
0 1
2 3
row-major -> 0 1 2 3
column-major -> 0 2 1 3
*/

/*
In = Coordenada n
Nn = Dimensión n
D = # de Dimensiones
IR = Índice Real Buscado

[n] -> [7] -> [3]
D1 -> IR = In -> IR = 3
[m][n] -> [3][3] -> [2][2]
D2 -> IR = (I0)(N1) + I1 -> (2)(2)+2 = 4 + 2 = 6
[m][n][p]
D3 -> D2N2+I2 -> [(I0)(N1)+I1]N2 + I2
*/
func rowMajor(dimension int, coordenadas []int, tamanio []int) int {
	var index int

	if dimension == 1 {
		index = coordenadas[0]
	} else {
		index = coordenadas[0]
		for i := 1; i < dimension; i++ {
			index = index*tamanio[i] + coordenadas[i]
		}
	}
	return index
}

func operar() {
	// [3][4] -> [2,2]
	/*
		1,1 | 1,2 | 1,3 | 1,4 | 2,1 | 2,2 | 2,3 | 2,4 | 3,1 | 3,2 | 3,3 | 3,4
		Loc (A[i1,i2]) = a + (i1-N1) * (M2-N2+1) + (i2-N2)
		IR = (I0)(N1) + I1
	*/
	dimension := 2
	coordenadas := []int{0, 2}
	tamanio := []int{3, 3}
	resultado := rowMajor(dimension, coordenadas, tamanio)
	fmt.Println(resultado)
}

func main() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("1. Inicio de Sesion")
		fmt.Println("2. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			MenuLogin()
		case 2:
			salir = true
		}
	}
}

func MenuLogin() {
	usuario := ""
	password := ""
	fmt.Print("Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_201700918" && password == "admin" {
		fmt.Println("Administrador Inicio Sesion")
		MenuAdmin()
	} else if listaDoble.Buscar(usuario, password) {
		fmt.Println("Bienvenido alumno ", usuario)
	} else {
		fmt.Println("ERROR EN CREDENCIALES!!!!")
	}
}

func MenuAdmin() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar de Cursos")
		fmt.Println("4. Control de Estudiantes")
		fmt.Println("5. Reportes")
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			CargaTutores()
		case 2:
			CargaEstudiantes()
		case 3:
			fmt.Println("Se cargo los cursos")
		case 4:
			ControlEstudiantes()
		case 5:
			fmt.Println("Mis Reportes")
		case 6:
			salir = true
		}

	}
}

func CargaTutores() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerCSV(ruta)
	fmt.Println("Se cargo a la Cola los tutores")
}

func CargaEstudiantes() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	listaDoble.LeerCSV(ruta)
	fmt.Println("Se cargo los estudiantes")
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		colaPrioridad.Primero_Cola()
		fmt.Println("════════════════════")
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			listaDobleCircular.Agregar(colaPrioridad.Primero.Tutor.Carnet, colaPrioridad.Primero.Tutor.Nombre, colaPrioridad.Primero.Tutor.Curso, colaPrioridad.Primero.Tutor.Nota)
			colaPrioridad.Descolar()
		} else if opcion == 2 {
			colaPrioridad.Descolar()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}
