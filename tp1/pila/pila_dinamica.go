package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	CAPACIDAD_INICIAL    = 10
	CAPACIDAD_PROPORCION = 2
	CANTIDAD_PROPORCION  = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	return pila
}

func (p *pilaDinamica[T]) redimensionar(capacidad_nueva int) {
	arreglo_nuevo := make([]T, capacidad_nueva)
	copy(arreglo_nuevo, p.datos)
	p.datos = arreglo_nuevo
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * CAPACIDAD_PROPORCION)
	}

	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	if p.cantidad*CANTIDAD_PROPORCION <= cap(p.datos) && CAPACIDAD_INICIAL < cap(p.datos) {
		p.redimensionar(cap(p.datos) / CAPACIDAD_PROPORCION)
	}

	p.cantidad--
	return p.datos[p.cantidad]
}
