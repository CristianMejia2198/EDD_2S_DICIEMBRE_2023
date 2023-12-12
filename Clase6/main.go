package main

import (
	"Clase6/estructuras/ColaPrioridad"
	"Clase6/estructuras/Listas"
	"Clase6/estructuras/MatrizDispersa"
	"fmt"
	"strconv"
)

var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var colaPrioridad *ColaPrioridad.Cola = &ColaPrioridad.Cola{Primero: nil, Longitud: 0}
var matrizDispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}

var loggeado_estudiante string = ""

func main() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Print("\033[H\033[2J")
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
	fmt.Print("\033[H\033[2J")
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
		loggeado_estudiante = usuario
		MenuEstudiantes()
	} else {
		fmt.Println("ERROR EN CREDENCIALES!!!!")
	}
}

func MenuAdmin() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar de Cursos") //Falta
		fmt.Println("4. Control de Estudiantes")
		fmt.Println("5. Reportes") //Falta
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

func MenuEstudiantes() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("1. Ver Tutores Disponibles")
		fmt.Println("2. Asignarse Tutores")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Print("\033[H\033[2J")
			listaDobleCircular.Mostrar()
		case 2:
			AsignarCurso()
		case 3:
			salir = true
		}
	}
}

func AsignarCurso() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Println("Teclee el codigo del curso: ")
		fmt.Scanln(&opcion)
		//Iria el primer If del Arbol (pendiente)
		if listaDobleCircular.Buscar(opcion) {
			TutorBuscado := listaDobleCircular.BuscarTutor(opcion)
			estudiante, err := strconv.Atoi(loggeado_estudiante)
			if err != nil {
				break
			}
			matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.Carnet, opcion)
			matrizDispersa.Reporte("Matriz.jpg")
			break
		} else {
			fmt.Println("No hay tutores para ese curso....")
			break
		}
	}
}

func CargaTutores() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerCSV(ruta)
	fmt.Println("Se cargo a la Cola los tutores")
}

func CargaEstudiantes() {
	fmt.Print("\033[H\033[2J")
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
		fmt.Print("\033[H\033[2J")
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
