package estructuras

import "fmt"

type Cola struct {
	Inicio   *NodoCola
	Longitud int
}

func (c *Cola) Encolar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) EncolarPrioridad(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	/*
		Ejemplo Lab
		Prioridad 1: 85 - 100
		Prioridad 2: 70 - 84
		Prioridad 3: 61 - 69

		Enunciado
		Prioridad 1: 90-100
		Prioridad 2: 75-89
		Prioridad 3: 65–74
		Prioridad 4: 64-61

	*/
	if nota >= 85 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 70 && nota <= 84 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 61 && nota <= 69 {
		nuevoNodo.Prioridad = 3
	} else {
		return
	}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad == nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad < nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else { // 1 -> 1 -> 3 Agregar 2
				aux = aux.Siguiente
			}
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay tutores en la cola")
	} else {
		c.Inicio = c.Inicio.Siguiente
		c.Longitud--
	}
}

func (c *Cola) Primero() {
	if c.Longitud == 0 {
		fmt.Println("No hay mas Tutores")
	} else {
		fmt.Println("Actual: ", c.Inicio.Tutor.Carnet)
		fmt.Println("Nombre: ", c.Inicio.Tutor.Nombre)
		fmt.Println("Curso: ", c.Inicio.Tutor.Curso)
		fmt.Println("Nota: ", c.Inicio.Tutor.Nota)
		fmt.Println("Prioridad: ", c.Inicio.Prioridad)
		if c.Inicio.Siguiente != nil {
			fmt.Println("Siguiente: ", c.Inicio.Siguiente.Tutor.Carnet)
		} else {
			fmt.Print("Siguiente: No hay mas tutores por evaluar")
		}
	}
}
