package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// --------------LISTA ENLAZADA--------------
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestListaInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.VerPrimero())
}

func TestListaInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 10, lista.VerUltimo())
	lista.InsertarUltimo(11)
	lista.InsertarUltimo(12)
	lista.InsertarUltimo(13)
	require.EqualValues(t, 13, lista.VerUltimo())
}

func TestListaBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	for i := 1; i < 4; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestListaVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.VerPrimero())
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	lista.InsertarUltimo(3)
	require.EqualValues(t, 1, lista.VerPrimero())
}

func TestListaVerUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerUltimo())
	lista.InsertarUltimo(3)
	require.EqualValues(t, 3, lista.VerUltimo())
	lista.InsertarPrimero(1)
	require.EqualValues(t, 3, lista.VerUltimo())
}

func TestListaControlDelLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.EqualValues(t, 0, lista.Largo())
	lista.InsertarPrimero(8)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(6)
	lista.InsertarUltimo(24)
	lista.InsertarUltimo(24)
	require.EqualValues(t, 5, lista.Largo())

	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.EqualValues(t, 2, lista.Largo())
}

func TestListaPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, 5, lista.Largo())
}

func TestListaVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 10000

	for i := 0; i < tam; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
	}

	for i := tam; i > 0; i-- {
		require.EqualValues(t, i-1, lista.BorrarPrimero())
	}
}

// --------------ITERADOR INTERNO--------------
func TestIteradorInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	var suma_ptr *int = &suma
	lista.Iterar(func(v int) bool {
		*suma_ptr += v
		return true
	})

	require.EqualValues(t, 0, suma)
}

func TestIteradorInternoRecorreTodaLaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	lista.InsertarUltimo(40)
	lista.InsertarUltimo(50)
	suma := 0
	var suma_ptr *int = &suma
	lista.Iterar(func(v int) bool {
		*suma_ptr += v
		return true
	})

	require.EqualValues(t, 150, suma)
}

func TestIteradorInternoCondicionCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("A")
	lista.InsertarUltimo("B")
	lista.InsertarUltimo("C")
	lista.InsertarUltimo("D")
	lista.InsertarUltimo("E")
	letras := ""
	var letras_ptr *string = &letras
	lista.Iterar(func(v string) bool {
		*letras_ptr += v
		return v != "C"
	})

	require.EqualValues(t, "ABC", letras)
}

// --------------ITERADOR EXTERNO--------------
func TestIteradorExternoInsertarAlPrincipioListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(0)
	require.EqualValues(t, 0, iter.VerActual())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())
}

func TestIteradorExternoInsertarAlPrincipioListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Insertar(0)

	iter2 := lista.Iterador()
	for i := 0; i < 4; i++ {
		require.True(t, iter2.HaySiguiente())
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}

	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 4, lista.Largo())
}

func TestIteradorExternoInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(4)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 4, lista.VerUltimo())
	require.EqualValues(t, 4, lista.Largo())
}

func TestIteradorExternoInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(2)

	iter2 := lista.Iterador()
	for i := 1; i <= 4; i++ {
		require.True(t, iter2.HaySiguiente())
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}

	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 4, lista.VerUltimo())
}

func TestIteradorExternoBorrarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorExternoBorrarAlPrincipioListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	valor := iter.Borrar()
	require.EqualValues(t, 1, valor)
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 2, iter.VerActual())
}

func TestIteradorExternoBorrarUltimoListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	for iter.VerActual() != lista.VerUltimo() {
		iter.Siguiente()
	}

	iter.Borrar()
	require.EqualValues(t, 4, lista.VerUltimo())
	require.EqualValues(t, 4, lista.Largo())
}

func TestIteradorExternoBorrarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual() == 3 {
			iter.Borrar()
		}
	}

	iter2 := lista.Iterador()
	for i := 1; i <= 5; i++ {
		require.True(t, iter2.HaySiguiente())
		if i == 3 {
			continue
		}
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}

	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 4, lista.Largo())
}

func TestIteradorExternoVerActual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	lista.InsertarPrimero(1)

	iter2 := lista.Iterador()
	require.EqualValues(t, 1, iter2.VerActual())
	require.True(t, iter2.HaySiguiente())
	iter2.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.VerActual() })
	require.False(t, iter2.HaySiguiente())
}

func TestIteradorExternoPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 1, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 2, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, lista.VerUltimo())
	require.False(t, iter.HaySiguiente())
	iter.Insertar(40)
	require.EqualValues(t, 40, lista.VerUltimo())
}

func TestIteradorExternoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	tam := 10000

	for i := 0; i < tam; i++ {
		iter.Insertar(i)
		require.EqualValues(t, i, iter.VerActual())
	}

	require.EqualValues(t, tam-1, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())

	iter2 := lista.Iterador()
	for i := tam; i > 0; i-- {
		require.EqualValues(t, i-1, iter2.Borrar())
	}
}
