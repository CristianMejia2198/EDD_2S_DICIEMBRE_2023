package main

import (
	"Clase14/estructuras"
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var GrafoGlobal = &estructuras.Grafo{Principal: nil}

type Cursos struct {
	Codigo        string   `json:"Codigo"`
	PostRequisito []string `json:"Post"`
}

type DatosCurso struct {
	Curso []Cursos `json:"Cursos"`
}

func GenerarGrafo(c *fiber.Ctx) error {
	var archivo estructuras.PeticionGrafo
	err := c.BodyParser(&archivo)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  400,
			"message": "Error al ingresar",
		})
	}
	LeerPeticion(archivo)
	return c.JSON(&fiber.Map{
		"status":  200,
		"message": "Lectura Existosa",
	})
}

func LeerPeticion(archivo estructuras.PeticionGrafo) {
	data, err := os.ReadFile(archivo.NombreArchivo)
	if err != nil {
		log.Fatal("Error al leer el json: ", err)
	}
	var datos DatosCurso
	err = json.Unmarshal(data, &datos)
	if err != nil {
		log.Fatal("Error al asignar el json: ", err)
	}
	for _, curso := range datos.Curso {
		if len(curso.PostRequisito) > 0 {
			for j := 0; j < len(curso.PostRequisito); j++ {
				GrafoGlobal.InsertarValores(curso.Codigo, curso.PostRequisito[j])
			}
		} else {
			GrafoGlobal.InsertarValores("ECYS", curso.Codigo)
		}
	}

	GrafoGlobal.Reporte("GrafoGlobal")
}

func main() {
	matriz := &estructuras.Grafo{Principal: &estructuras.NodoListaAdyacencia{Valor: "ECYS"}}
	matriz1 := &estructuras.Grafo{Principal: nil}
	Estatico(matriz)
	Lectura(matriz1)
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/grafo", GenerarGrafo)

	app.Listen(":4000")

}

func Estatico(matriz *estructuras.Grafo) {
	matriz.InsertarValores("ECYS", "0770")
	matriz.InsertarValores("ECYS", "0960")

	matriz.InsertarValores("0770", "0771")
	matriz.InsertarValores("0770", "0796")
	matriz.InsertarValores("0770", "0962")

	matriz.InsertarValores("0771", "0772")
	matriz.InsertarValores("0771", "0964")
	matriz.InsertarValores("0771", "0777")

	matriz.InsertarValores("0772", "0722")
	matriz.InsertarValores("0772", "0781")
	matriz.InsertarValores("0777", "0781")
	matriz.Reporte("Grafo1")
}

func Lectura(matriz1 *estructuras.Grafo) {
	data, err := os.ReadFile("cursos.json")
	if err != nil {
		log.Fatal("Error al leer el json: ", err)
	}
	var datos DatosCurso
	err = json.Unmarshal(data, &datos)
	if err != nil {
		log.Fatal("Error al asignar el json: ", err)
	}
	for _, curso := range datos.Curso {
		if len(curso.PostRequisito) > 0 {
			for j := 0; j < len(curso.PostRequisito); j++ {
				matriz1.InsertarValores(curso.Codigo, curso.PostRequisito[j])
			}
		} else {
			matriz1.InsertarValores("ECYS", curso.Codigo)
		}
	}

	matriz1.Reporte("Grafo2")
}
