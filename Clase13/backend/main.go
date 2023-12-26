package main

import (
	"Clase13/estructuras"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Validar(c *fiber.Ctx) error {
	var usuario estructuras.PeticionLogin
	c.BodyParser(&usuario)
	if usuario.UserName == "ADMIN_201700918" {
		if usuario.Passwrod == "admin" {
			return c.JSON(&fiber.Map{
				"status":  200,
				"message": "Credenciales Correctas",
			})
		}
	} else {
		//Validar Carnet
	}
	return c.JSON(&fiber.Map{
		"status":  400,
		"message": "Credenciales incorrectas",
	})
}

func main() {
	tablitaHash := estructuras.TablaHash{Tabla: make(map[int]estructuras.NodoHash), Capacidad: 7, Utilizacion: 0}
	/*tablitaHash.Insertar(201700918, "Cristian", "123")
	tablitaHash.Insertar(201712345, "Cristian", "123")
	tablitaHash.Insertar(201789369, "Cristian", "123")
	tablitaHash.Insertar(201700236, "Cristian", "123")
	tablitaHash.Insertar(201759369, "Cristian", "123")*/
	tablitaHash.LeerCSV("EstudiantesF1.csv")
	fmt.Println("Capacidad: ", tablitaHash.Capacidad)
	for i := 0; i < tablitaHash.Capacidad; i++ {
		if usuario, existe := tablitaHash.Tabla[i]; existe {
			fmt.Println("Posicion: ", i, " Carnet: ", usuario.Persona.Carnet)
		}
	} // 0,1,1,2,3,5,8,13,21,34

	fmt.Println("---------------------------------------------------")
	for i := 0; i < tablitaHash.Capacidad; i++ {
		if usuario, existe := tablitaHash.Tabla[i]; existe {
			fmt.Println("Llave: ", usuario.Llave, " Carnet: ", usuario.Persona.Carnet)
		}
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login-server", Validar)

	app.Listen(":4000")
}
