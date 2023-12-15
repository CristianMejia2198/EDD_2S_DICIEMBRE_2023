package ArbolAVL

type NodoArbol struct {
	Izquierdo         *NodoArbol
	Derecho           *NodoArbol
	Valor             string
	Altura            int
	Factor_Equilibrio int
}
