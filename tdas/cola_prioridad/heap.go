package cola_prioridad

const (
	CAPACIDAD_INICIAL  = 10
	FACTOR_REDIMENSION = 2
	FACTOR_ACHICAR     = FACTOR_REDIMENSION * FACTOR_REDIMENSION
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap(funcion_cmp)
	}

	heap := new(heap[T])
	heap.datos = make([]T, len(arreglo))
	copy(heap.datos, arreglo)
	heapify(heap.datos, funcion_cmp)
	heap.cmp = funcion_cmp
	heap.cantidad = len(arreglo)
	return heap
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == cap(h.datos) {
		h.redimensionar(cap(h.datos) * FACTOR_REDIMENSION)
	}

	h.datos[h.cantidad] = elem
	upHeap(h.datos, h.cantidad, h.cmp)
	h.cantidad++
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	dato := h.datos[0]
	swap(&h.datos[0], &h.datos[h.cantidad-1])
	h.cantidad--
	downHeap(h.datos, h.cantidad, 0, h.cmp)

	if h.cantidad*FACTOR_ACHICAR <= cap(h.datos) && CAPACIDAD_INICIAL < cap(h.datos) {
		h.redimensionar(cap(h.datos) / FACTOR_REDIMENSION)
	}

	return dato
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	heapsort(elementos, funcion_cmp, len(elementos))
}

func heapsort[T any](elementos []T, funcion_cmp func(T, T) int, largo int) {
	if largo <= 1 {
		return
	}

	swap(&elementos[0], &elementos[largo-1])
	downHeap(elementos, largo-1, 0, funcion_cmp)
	heapsort(elementos, funcion_cmp, largo-1)
}

func (h *heap[T]) redimensionar(capacidad_nueva int) {
	arreglo_nuevo := make([]T, capacidad_nueva)
	copy(arreglo_nuevo, h.datos)
	h.datos = arreglo_nuevo
}

func max[T any](arreglo []T, cmp func(T, T) int, largo, hIzq, hDer int) int {
	if hIzq == largo-1 || cmp(arreglo[hIzq], arreglo[hDer]) >= 0 {
		return hIzq
	}

	return hDer
}

func upHeap[T any](arreglo []T, posicionHijo int, cmp func(T, T) int) {
	if posicionHijo == 0 {
		return
	}

	posicionPadre := (posicionHijo - 1) / 2
	if cmp(arreglo[posicionPadre], arreglo[posicionHijo]) <= 0 {
		swap(&arreglo[posicionPadre], &arreglo[posicionHijo])
		upHeap(arreglo, posicionPadre, cmp)
	}
}

func downHeap[T any](arreglo []T, largo, posicionPadre int, cmp func(T, T) int) {
	hIzq := 2*posicionPadre + 1
	hDer := 2*posicionPadre + 2

	if hIzq >= largo {
		return
	}

	maximo := max(arreglo, cmp, largo, hIzq, hDer)

	if cmp(arreglo[posicionPadre], arreglo[maximo]) < 0 {
		swap(&arreglo[posicionPadre], &arreglo[maximo])
		downHeap(arreglo, largo, maximo, cmp)
	}
}

func heapify[T any](elementos []T, funcion_cmp func(T, T) int) {
	for i := len(elementos) - 1; i >= 0; i-- {
		downHeap(elementos, len(elementos), i, funcion_cmp)
	}
}

func swap[T any](x *T, y *T) {
	*x, *y = *y, *x
}
