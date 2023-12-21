package main

import (
	"Clase11/estructuras"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var arbolitoB *estructuras.ArbolB = &estructuras.ArbolB{Raiz: nil, Orden: 3}

func AgregarB(c *fiber.Ctx) error {
	fmt.Println("Recibi Solicitud")
	var nuevoElemento estructuras.SolicitudArbolB
	err := c.BodyParser(&nuevoElemento)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  400,
			"message": "Error al ingresar ",
		})
	}
	arbolitoB.Insertar(nuevoElemento.Valor)
	return c.JSON(&fiber.Map{
		"status":  200,
		"message": "Valor ingresado",
	})
}

func GenerarReporte(c *fiber.Ctx) error {
	fmt.Println("Recibi solicitud")
	var reporte estructuras.SolicitudReporte
	c.BodyParser(&reporte)
	if reporte.Estructura_solicitada == "Arbol B" {
		arbolitoB.Graficar(reporte.Nombre)
		return c.JSON(&fiber.Map{
			"status":  200,
			"message": "Grafica Generada",
		})
	} else {
		return c.JSON(&fiber.Map{
			"status":  400,
			"message": "Error",
		})
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	//arbolitoB.Buscar(17)
	app.Post("/ingresar-arbol-b", AgregarB)
	app.Post("/reporte", GenerarReporte)
	app.Listen(":4000")
}
