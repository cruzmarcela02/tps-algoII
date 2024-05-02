package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el elemento al principio de la lista, si la lista está vacía
	// el elemento se inserta al final.
	InsertarPrimero(T)

	// InsertarUltimo inserta el elemento al final de la lista.
	// Si la lista está vacía, el insertado será el primero de la lista.
	InsertarUltimo(T)

	// BorrarPrimero, borra el primer elemento de la lista.
	// Si la lista está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. Si está vacía, entra en pánico
	// con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve el largo de la lista.
	Largo() int

	// Iterar recibe una función por parámetro que aplica a todos los elementos de la lista o hasta
	// que la función devuelva false.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un IteradorLista de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el valor del elemento en donde está parado el iterador. Si no hay siguiente,
	// entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve verdadero si obtiene el elemento en donde está parado el iterador,
	// false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemento. Si no hay siguiente, entra en pánico con
	// un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar inserta el elemento entre el elemento anterior y el actual,
	// manteniendo el anterior y actualizando el actual al elemento insertado.
	Insertar(T)

	// Borrar borra el elemento actual de la lista, y lo devuelve.
	Borrar() T
}
