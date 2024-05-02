package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodoNuevo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.ultimo = nodoNuevo
	} else {
		nodoNuevo.prox = lista.primero
	}

	lista.primero = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodoNuevo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = nodoNuevo
	} else {
		lista.ultimo.prox = nodoNuevo
	}

	lista.ultimo = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.prox

	if lista.primero == nil {
		lista.ultimo = nil
	}

	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// ITERADOR INTERNO
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}

		actual = actual.prox
	}
}

// ITERADOR EXTERNO
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iterListaEnlazada[T])
	iterador.actual = lista.primero
	iterador.lista = lista
	return iterador
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.prox
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)

	if iterador.lista.EstaVacia() {
		iterador.lista.primero = nuevoNodo
		iterador.lista.ultimo = nuevoNodo
	} else if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
		nuevoNodo.prox = iterador.actual
	} else {
		iterador.anterior.prox = nuevoNodo
		nuevoNodo.prox = iterador.actual
		if !iterador.HaySiguiente() {
			iterador.lista.ultimo = nuevoNodo
		}
	}

	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	dato := iterador.VerActual()

	if iterador.anterior == nil {
		iterador.lista.primero = iterador.lista.primero.prox
		iterador.actual = iterador.lista.primero
		if iterador.lista.primero == nil {
			iterador.lista.ultimo = nil
		}
	} else {
		iterador.anterior.prox = iterador.actual.prox
		iterador.actual = iterador.actual.prox
		if iterador.actual == nil {
			iterador.lista.ultimo = iterador.anterior
		}
	}

	iterador.lista.largo--
	return dato
}
