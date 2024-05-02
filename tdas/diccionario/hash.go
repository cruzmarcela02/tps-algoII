package diccionario

import "fmt"

type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)

const (
	CAPACIDAD_INICIAL   = 17
	MAXIMA_CARGA        = 0.7
	MINIMA_CARGA        = 0.2
	CANTIDAD_PROPORCION = 2
)

type campo[K comparable, V any] struct {
	clave  K
	valor  V
	estado estado
}

type hash[K comparable, V any] struct {
	tabla     []campo[K, V]
	cantidad  int
	capacidad int
	borrados  int
}

type iterDiccionario[K comparable, V any] struct {
	posicion int
	hash     *hash[K, V]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hash[K, V])
	hash.tabla = crearTabla[K, V](CAPACIDAD_INICIAL)
	hash.capacidad = CAPACIDAD_INICIAL
	return hash
}

func crearTabla[K comparable, V any](capacidadNueva int) []campo[K, V] {
	tabla := make([]campo[K, V], capacidadNueva)
	for i := range tabla {
		tabla[i].estado = VACIO
	}

	return tabla
}

func (h *hash[K, V]) Guardar(clave K, dato V) {
	if float64(h.cantidad+h.borrados) >= float64(h.capacidad)*MAXIMA_CARGA {
		h.redimensionar(h.capacidad * CANTIDAD_PROPORCION)
	}

	posicion := h.buscar(clave)
	if h.tabla[posicion].estado == VACIO {
		h.tabla[posicion].clave = clave
		h.tabla[posicion].estado = OCUPADO
		h.cantidad++
	}

	h.tabla[posicion].valor = dato
}

func (h *hash[K, V]) Pertenece(clave K) bool {
	posicion := h.buscar(clave)
	return h.tabla[posicion].estado != VACIO
}

func (h *hash[K, V]) buscar(clave K) int {
	posicion := hashing(clave, h.capacidad)

	for h.tabla[posicion].estado != VACIO {
		if h.tabla[posicion].estado == OCUPADO && h.tabla[posicion].clave == clave {
			return posicion
		}

		if posicion == h.capacidad-1 {
			posicion = 0
		} else {
			posicion++
		}
	}

	return posicion
}

func (h *hash[K, V]) Obtener(clave K) V {
	posicion := h.buscar(clave)
	if h.tabla[posicion].estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}

	return h.tabla[posicion].valor
}

func (h *hash[K, V]) Borrar(clave K) V {
	posicion := h.buscar(clave)
	if h.tabla[posicion].estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}

	valor := h.tabla[posicion].valor
	h.tabla[posicion].estado = BORRADO
	h.cantidad--
	h.borrados++

	if float64(h.cantidad+h.borrados) <= float64(h.capacidad)*MINIMA_CARGA && CAPACIDAD_INICIAL < h.capacidad {
		h.redimensionar(h.capacidad / CANTIDAD_PROPORCION)
	}

	return valor
}

func (h *hash[K, V]) Cantidad() int {
	return h.cantidad
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hash[K, V]) redimensionar(capacidadNueva int) {
	tablaNueva := crearTabla[K, V](capacidadNueva)
	tablaVieja := h.tabla
	h.tabla = tablaNueva
	h.capacidad = capacidadNueva
	h.cantidad = 0
	h.borrados = 0

	for _, elemento := range tablaVieja {
		if elemento.estado == OCUPADO {
			h.Guardar(elemento.clave, elemento.valor)
		}
	}
}

// https://blog.fredrb.com/2021/04/01/hashtable-go/
func hashing[K comparable](clave K, capacidad int) int {
	var hash, fnvPrime uint64 = 14695981039346656037, 1099511628211

	for _, b := range convertirABytes(clave) {
		hash = hash ^ uint64(b)
		hash = hash * fnvPrime
	}

	return int(hash % uint64(capacidad))
}

// ITERADOR INTERNO
func (h *hash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, campo := range h.tabla {
		if campo.estado == OCUPADO && !visitar(campo.clave, campo.valor) {
			break
		}
	}
}

// ITERADOR EXTERNO
func (h *hash[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(iterDiccionario[K, V])
	iterador.hash = h
	actualizarPosicion(iterador)
	return iterador
}

func (iterador *iterDiccionario[K, V]) HaySiguiente() bool {
	return iterador.posicion < iterador.hash.capacidad
}

func (iterador *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iterador.hash.tabla[iterador.posicion].clave, iterador.hash.tabla[iterador.posicion].valor
}

func (iterador *iterDiccionario[K, V]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iterador.posicion++
	actualizarPosicion(iterador)
}

func actualizarPosicion[K comparable, V any](iterador *iterDiccionario[K, V]) {
	for iterador.posicion < iterador.hash.capacidad {
		if iterador.hash.tabla[iterador.posicion].estado == OCUPADO {
			return
		}

		iterador.posicion++
	}

	iterador.posicion = iterador.hash.capacidad
}
