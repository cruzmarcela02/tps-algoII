package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaPocosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.VerTope())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.VerTope())
	require.EqualValues(t, 1, pila.Desapilar())
}

func TestPilaVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	tam := 10000

	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}

	for i := tam; i > 0; i-- {
		require.EqualValues(t, i-1, pila.Desapilar())
	}
}

func TestPilaDesapilarHastaQueEsteVacia(t *testing.T) {
	tam := 10
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < tam; i++ {
		pila.Apilar(i)
	}

	for i := tam; i > 0; i-- {
		require.EqualValues(t, i-1, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
}

func TestPilaDesapilarInvalido(t *testing.T) {
	tam := 10
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < tam; i++ {
		pila.Apilar(i)
	}

	for i := tam; i > 0; i-- {
		require.EqualValues(t, i-1, pila.Desapilar())
	}

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("A")
	pila.Apilar("B")
	pila.Apilar("C")
	require.EqualValues(t, "C", pila.VerTope())
	require.EqualValues(t, "C", pila.Desapilar())
	require.EqualValues(t, "B", pila.Desapilar())
	require.EqualValues(t, "A", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

func TestPilaFloat64(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(4.5)
	pila.Apilar(80.1)
	pila.Apilar(10000.5)
	require.EqualValues(t, 10000.5, pila.VerTope())
	require.EqualValues(t, 10000.5, pila.Desapilar())
	require.EqualValues(t, 80.1, pila.Desapilar())
	require.EqualValues(t, 4.5, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}
