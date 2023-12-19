package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Persona struct {
	Nombre string `json:"Nombre"`
	Edad   int    `json:"Edad"`
}

/*
func NombreDeFuncion(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

*/

func Principal(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"status":  200,
		"message": "Funcion 2",
	})
}

func Validar(c *fiber.Ctx) error {
	var persona Persona
	/*
		{
			"Nombre": "hola"
			"Edad": 25
		}
	*/
	c.BodyParser(&persona)
	fmt.Println(persona)
	if persona.Edad >= 18 {
		return c.JSON(&fiber.Map{
			"status":  200,
			"message": "Usted es mayor de edad",
			"nombre":  persona.Nombre,
		})
	} else {
		return c.JSON(&fiber.Map{
			"status":  200,
			"message": "Usted es menor de edad",
		})
	}
}

func main() {
	app := fiber.New()

	//GET (llamar a la un ruta del servidor y que nos devuelva una respuesta)
	//POST (Recibe parametros, proceso la informacion y devuelve una respuesta)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status":  200,
			"message": "Funcion 1",
		})
	})

	app.Get("/principal", Principal)
	app.Post("/validar-edad", Validar)

	app.Listen(":4000")
}
