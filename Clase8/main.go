package main

import (
	"Clase8/estructuras"
	"encoding/json"
	"log"
	"os"
)

type Curso struct {
	Codigo string `json:"Codigo"`
	Nombre string `json:"Nombre"`
}

type DatosCursos struct {
	Cursos []Curso `json:"Cursos"`
}

func main() {
	arbolavl := estructuras.ArbolAVL{Raiz: nil}
	/*arbolavl.InsertarElemento("0770")
	arbolavl.InsertarElemento("0771")
	arbolavl.InsertarElemento("0772")
	arbolavl.InsertarElemento("0773")
	arbolavl.InsertarElemento("0774")
	arbolavl.InsertarElemento("0775")
	arbolavl.InsertarElemento("0777")*/
	//arbolavl.Graficar()

	/*Lectura de Json*/
	data, err := os.ReadFile("archivo.json")
	if err != nil {
		log.Fatal("Error al leer el archivo:", err)
	}

	var datos DatosCursos
	err = json.Unmarshal(data, &datos)
	if err != nil {
		log.Fatal("Error al decodificar el JSON:", err)
	}
	for _, curso := range datos.Cursos {
		arbolavl.InsertarElemento(curso.Codigo)
	}
	arbolavl.Graficar()
}
