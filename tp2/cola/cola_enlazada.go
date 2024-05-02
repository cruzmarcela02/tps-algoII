package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevo := crearNodo(dato)

	if cola.EstaVacia() {
		cola.primero = nuevo
	} else {
		cola.ultimo.prox = nuevo
	}

	cola.ultimo = nuevo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	dato := cola.primero.dato
	cola.primero = cola.primero.prox

	if cola.primero == nil {
		cola.ultimo = nil
	}

	return dato
}
