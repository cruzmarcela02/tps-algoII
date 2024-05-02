package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaPocosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(2)
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.EqualValues(t, 1, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.VerPrimero())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.VerPrimero())
	require.EqualValues(t, 3, cola.Desencolar())
}

func TestColaVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	tam := 10000

	for i := 0; i < tam; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}

	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
}

func TestColaDesencolarHastaQueEsteVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	tam := 10

	for i := 0; i < tam; i++ {
		cola.Encolar(i)
	}

	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())
}

func TestColaDesencolarInvalido(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	tam := 10

	for i := 0; i < tam; i++ {
		cola.Encolar(i)
	}

	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("A")
	cola.Encolar("B")
	cola.Encolar("C")
	require.EqualValues(t, "A", cola.VerPrimero())
	require.EqualValues(t, "A", cola.Desencolar())
	require.EqualValues(t, "B", cola.Desencolar())
	require.EqualValues(t, "C", cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestColaFloat64(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	cola.Encolar(4.5)
	cola.Encolar(83.1)
	cola.Encolar(11000.5)
	require.EqualValues(t, 4.5, cola.VerPrimero())
	require.EqualValues(t, 4.5, cola.Desencolar())
	require.EqualValues(t, 83.1, cola.Desencolar())
	require.EqualValues(t, 11000.5, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}
