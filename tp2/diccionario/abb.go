package diccionario

import Pila "algueiza/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iteradorAbb[K comparable, V any] struct {
	abb     *abb[K, V]
	desde   *K
	hasta   *K
	pilaIzq Pila.Pila[*nodoAbb[K, V]]
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodoAbb := new(nodoAbb[K, V])
	nodoAbb.clave = clave
	nodoAbb.dato = dato
	return nodoAbb
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dicc := new(abb[K, V])
	dicc.cmp = funcion_cmp
	return dicc
}

func (abb *abb[K, V]) buscarNodo(clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	return buscar(abb.raiz, abb.raiz, clave, abb.cmp)
}

func buscar[K comparable, V any](actual, anterior *nodoAbb[K, V], clave K, cmp func(K, K) int) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if actual == nil {
		return actual, anterior
	}

	numero := cmp(clave, actual.clave)
	if numero == 0 {
		return actual, anterior
	}

	if numero > 0 {
		return buscar(actual.derecho, actual, clave, cmp)
	} else {
		return buscar(actual.izquierdo, actual, clave, cmp)
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	actual, _ := abb.buscarNodo(clave)
	return actual != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	actual, _ := abb.buscarNodo(clave)
	if actual == nil {
		panic("La clave no pertenece al diccionario")
	}

	return actual.dato
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo := crearNodoABB(clave, dato)
	if abb.raiz == nil {
		abb.raiz = nodo
		abb.cantidad++
		return
	}

	actual, anterior := abb.buscarNodo(clave)
	if actual == nil {
		if abb.cmp(clave, anterior.clave) < 0 {
			anterior.izquierdo = nodo
		} else {
			anterior.derecho = nodo
		}

		abb.cantidad++
	} else {
		actual.dato = nodo.dato
	}
}

func (abb *abb[K, V]) Borrar(clave K) V {
	actual, anterior := abb.buscarNodo(clave)
	if actual == nil {
		panic("La clave no pertenece al diccionario")
	}

	dato := actual.dato
	numero := abb.cmp(actual.clave, anterior.clave)

	if !tieneDoshijos(actual) {
		abb.borrarUnoCeroHijos(numero, actual, anterior)
		abb.cantidad--
	} else {
		reemplazo := buscarReemplazante(actual.derecho)
		abb.Borrar(reemplazo.clave)
		actual.clave = reemplazo.clave
		actual.dato = reemplazo.dato
	}

	return dato
}

func (abb *abb[K, V]) borrarUnoCeroHijos(numero int, actual, anterior *nodoAbb[K, V]) {
	if numero < 0 {
		if actual.izquierdo == nil {
			anterior.izquierdo = actual.derecho
		} else {
			anterior.izquierdo = actual.izquierdo
		}
	} else if numero > 0 {
		if actual.izquierdo == nil {
			anterior.derecho = actual.derecho
		} else {
			anterior.derecho = actual.izquierdo
		}
	} else {
		if actual.izquierdo == nil {
			abb.raiz = actual.derecho
		} else {
			abb.raiz = actual.izquierdo
		}
	}
}

func tieneDoshijos[K comparable, V any](actual *nodoAbb[K, V]) bool {
	return actual.izquierdo != nil && actual.derecho != nil
}

func buscarReemplazante[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}

	if nodo.izquierdo == nil {
		return nodo
	}

	return buscarReemplazante(nodo.izquierdo)
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

// ITERADOR INTERNO
func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRango(nil, nil, visitar, abb.cmp)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || cmp(nodo.clave, *desde) > 0 {
		if !nodo.izquierdo.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if hasta == nil || cmp(nodo.clave, *hasta) < 0 {
		if !nodo.derecho.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	return true
}

// ITERADOR EXTERNO
func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := new(iteradorAbb[K, V])
	iterador.abb = abb
	iterador.desde = desde
	iterador.hasta = hasta
	iterador.pilaIzq = Pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador.apilarIzq(iterador.abb.raiz)
	return iterador
}

func (iterador *iteradorAbb[K, V]) apilarIzq(actual *nodoAbb[K, V]) {
	if actual == nil {
		return
	}

	if iterador.desde != nil && iterador.abb.cmp(actual.clave, *iterador.desde) < 0 {
		iterador.apilarIzq(actual.derecho)
	}

	if (iterador.desde == nil || iterador.abb.cmp(actual.clave, *iterador.desde) >= 0) && (iterador.hasta == nil || iterador.abb.cmp(actual.clave, *iterador.hasta) <= 0) {
		iterador.pilaIzq.Apilar(actual)
		iterador.apilarIzq(actual.izquierdo)
	}

	if iterador.hasta != nil && iterador.abb.cmp(actual.clave, *iterador.hasta) > 0 {
		iterador.apilarIzq(actual.izquierdo)
	}
}

func (iterador *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iterador.pilaIzq.EstaVacia()
}

func (iterador *iteradorAbb[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := iterador.pilaIzq.VerTope()
	return actual.clave, actual.dato
}

func (iterador *iteradorAbb[K, V]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	desapilado := iterador.pilaIzq.Desapilar()
	iterador.apilarIzq(desapilado.derecho)
}
