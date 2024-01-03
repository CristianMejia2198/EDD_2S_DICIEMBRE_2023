package tablaHash

import "strconv"

type TablaHash struct {
	Tabla       map[int]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) calculoIndice(carnet int) int {
	var numeros []int
	for {
		if carnet > 0 {
			digito := carnet % 10
			numeros = append([]int{digito}, numeros...)
			carnet = carnet / 10
		} else {
			break
		}
	}

	var numeros_ascii []rune
	for _, numero := range numeros {
		valor := rune(numero + 48)
		numeros_ascii = append(numeros_ascii, valor)
	}

	final := 0
	for _, numero_ascii := range numeros_ascii {
		final += int(numero_ascii)
	}

	indice := final % t.Capacidad
	return indice
}

func (t *TablaHash) capacidadTabla() {
	auxCap := float64(t.Capacidad) * 0.6
	if t.Utilizacion > int(auxCap) {
		auxAnterior := t.Capacidad
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar(auxAnterior)
	}
}
func (t *TablaHash) nuevaCapacidad() int {
	contador := 0
	a, b := 0, 1
	for contador < 100 {
		contador++
		if a > t.Capacidad {
			return a
		}
		a, b = b, a+b
	}
	return a
}

func (t *TablaHash) reInsertar(capacidadAnterior int) {
	auxTabla := t.Tabla
	t.Tabla = make(map[int]NodoHash)
	for i := 0; i < capacidadAnterior; i++ {
		if usuario, existe := auxTabla[i]; existe {
			t.Insertar(usuario.Persona.Carnet, usuario.Persona.Nombre, usuario.Persona.Password, usuario.Persona.Cursos)
		}
	}
}

func (t *TablaHash) reCalculoIndice(carnet int, contador int) int {
	nuevoIndice := t.calculoIndice(carnet) + (contador * contador) //5+4=9
	return t.nuevoIndice(nuevoIndice)
}

func (t *TablaHash) nuevoIndice(nuevoIndice int) int {
	nuevoPosicion := 0
	if nuevoIndice < t.Capacidad {
		nuevoPosicion = nuevoIndice
	} else {
		nuevoPosicion = nuevoIndice - t.Capacidad
		nuevoPosicion = t.nuevoIndice(nuevoPosicion)
	}
	return nuevoPosicion
}

func (t *TablaHash) Insertar(carnet int, nombre string, password string, cursos []string) { // cursos []string
	indice := t.calculoIndice(carnet)
	nuevoNodo := &NodoHash{Llave: indice, Persona: &Persona{Carnet: carnet, Nombre: nombre, Password: password, Cursos: cursos}}
	if indice < t.Capacidad {
		if _, existe := t.Tabla[indice]; !existe {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidadTabla()
		} else {
			contador := 1
			indice = t.reCalculoIndice(carnet, contador)
			for {
				if _, existe := t.Tabla[indice]; existe {
					contador++
					indice = t.reCalculoIndice(carnet, contador)
				} else {
					nuevoNodo.Llave = indice
					t.Tabla[indice] = *nuevoNodo
					t.Utilizacion++
					t.capacidadTabla()
					break
				}
			}
		}
	}
}

func (t *TablaHash) Buscar(carnet string, password string) bool {
	valTemp, err := strconv.Atoi(carnet)
	if err != nil {
		return false
	}
	indice := t.calculoIndice(valTemp)
	if indice < t.Capacidad {
		if usuario, existe := t.Tabla[indice]; existe {
			if usuario.Persona.Carnet == valTemp {
				if usuario.Persona.Password == password {
					return true
				}
			} else {
				contador := 1
				indice = t.reCalculoIndice(valTemp, contador)
				for {
					if usuario, existe := t.Tabla[indice]; existe {
						if usuario.Persona.Carnet == valTemp {
							if usuario.Persona.Password == password {
								return true
							} else {
								return false
							}
						} else {
							contador++
							indice = t.reCalculoIndice(valTemp, contador)
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}

func (t *TablaHash) ConvertirArreglo() []NodoHash {
	var arrays []NodoHash
	if t.Utilizacion > 0 {
		for i := 0; i < t.Capacidad; i++ {
			if usuario, existe := t.Tabla[i]; existe {
				arrays = append(arrays, usuario)
			}
		}
	}
	return arrays
}

func (t *TablaHash) BuscarSesion(carnet string) *Persona {
	valTemp, err := strconv.Atoi(carnet)
	if err != nil {
		return nil
	}
	indice := t.calculoIndice(valTemp)
	if indice < t.Capacidad {
		if usuario, existe := t.Tabla[indice]; existe {
			if usuario.Persona.Carnet == valTemp {
				return usuario.Persona
			} else {
				contador := 1
				indice = t.reCalculoIndice(valTemp, contador)
				for {
					if usuario, existe := t.Tabla[indice]; existe {
						if usuario.Persona.Carnet == valTemp {
							return usuario.Persona
						} else {
							contador++
							indice = t.reCalculoIndice(valTemp, contador)
						}
					} else {
						return nil
					}
				}
			}
		}
	}
	return nil
}
