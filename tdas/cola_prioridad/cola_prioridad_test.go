package cola_prioridad_test

import (
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmpEnteros(clave1, clave2 int) int {
	return clave1 - clave2
}

func TestHeapVacio(t *testing.T) {
	t.Log("Comprueba que el heap no tiene elementos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapActualizaCantidad(t *testing.T) {
	t.Log("Comprueba que el maximo del heap se actualice")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(10)
	heap.Encolar(11)
	heap.Encolar(5)
	require.EqualValues(t, 3, heap.Cantidad())

	heap.Desencolar()
	heap.Desencolar()
	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
}

func TestHeapActualizaMaximos(t *testing.T) {
	t.Log("Comprueba que el maximo del heap se actualice al encolar varios elementos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(3)
	require.EqualValues(t, 3, heap.VerMax())
	heap.Encolar(8)
	require.EqualValues(t, 8, heap.VerMax())
	heap.Encolar(5)
	require.EqualValues(t, 8, heap.VerMax())

	require.EqualValues(t, 3, heap.Cantidad())

	heap.Encolar(10)
	require.EqualValues(t, 10, heap.VerMax())

	require.False(t, heap.EstaVacia())
}

func TestHeapActualizaMaximosAlDesencolar(t *testing.T) {
	t.Log("Comprueba que el maximo del heap se actualice al desencolar varios elementos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(30)
	heap.Encolar(7)
	heap.Encolar(15)
	heap.Encolar(27)
	require.EqualValues(t, 30, heap.VerMax())

	require.EqualValues(t, 30, heap.Desencolar())
	require.EqualValues(t, 27, heap.Desencolar())
	require.EqualValues(t, 15, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())

	require.False(t, heap.EstaVacia())
}

func TestHeapActualizaMaximosStrings(t *testing.T) {
	t.Log("Comprueba que el maximo del heap se actualice al encolar varios elementos")
	heap := TDAHeap.CrearHeap(strings.Compare)
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar("Banana")
	require.EqualValues(t, "Banana", heap.VerMax())
	heap.Encolar("Frutilla")
	require.EqualValues(t, "Frutilla", heap.VerMax())
	heap.Encolar("Anana")
	require.EqualValues(t, "Frutilla", heap.VerMax())

	require.EqualValues(t, 3, heap.Cantidad())

	heap.Encolar("Zanahoria")
	require.EqualValues(t, "Zanahoria", heap.VerMax())

	require.False(t, heap.EstaVacia())
}

func TestHeapActualizaMaximosAlDesencolarStrings(t *testing.T) {
	t.Log("Comprueba que el maximo del heap se actualice al desencolar varios elementos")
	heap := TDAHeap.CrearHeap(strings.Compare)
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar("Gato")
	heap.Encolar("Perro")
	heap.Encolar("Tortuga")
	heap.Encolar("Delfin")
	require.EqualValues(t, "Tortuga", heap.VerMax())
	require.EqualValues(t, 4, heap.Cantidad())

	require.EqualValues(t, "Tortuga", heap.Desencolar())
	require.EqualValues(t, "Perro", heap.Desencolar())
	require.EqualValues(t, "Gato", heap.Desencolar())
	require.EqualValues(t, "Delfin", heap.VerMax())

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
}

func TestHeapVolumen(t *testing.T) {
	t.Log("Comprueba el heap con muchos elementos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	tam := 10000

	for i := 0; i <= tam; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax())
	}

	for i := tam; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
}

func TestHeapConArregloDado(t *testing.T) {
	t.Log("Comprueba que el heap con un slice dado, no vacio")
	slice := []int{4, 2, 7, 10, 4, 3, 8, 5}
	heap := TDAHeap.CrearHeapArr(slice, cmpEnteros)

	require.EqualValues(t, len(slice), heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 10, heap.Desencolar())
}

func TestHeapConArregloDadoNoCambiaOriginal(t *testing.T) {
	t.Log("Crea un heap con un slice dado y no cambia el slice original")
	slice := []int{4, 2, 7, 10, 4, 3, 8, 5}
	heap := TDAHeap.CrearHeapArr(slice, cmpEnteros)

	require.EqualValues(t, len(slice), heap.Cantidad())
	require.False(t, heap.EstaVacia())
	heap.Desencolar()
	heap.Desencolar()

	require.Equal(t, []int{4, 2, 7, 10, 4, 3, 8, 5}, slice, "Slice original")
}

func TestHeapConArregloUnElemento(t *testing.T) {
	t.Log("Crea un heap con un slice de un elemento dado")
	slice := []int{10}
	heap := TDAHeap.CrearHeapArr(slice, cmpEnteros)

	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 10, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapConArregloDadoVacio(t *testing.T) {
	t.Log("Comprueba que el heap con un slice dado, vacio")
	slice := []int{}
	heap := TDAHeap.CrearHeapArr(slice, cmpEnteros)

	require.EqualValues(t, len(slice), heap.Cantidad())
	require.True(t, heap.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapConArregloDadoVacioYConElementos(t *testing.T) {
	t.Log("Comprueba que crear un heap desde un arreglo vacio no rompe luego al agregar elementos")
	slice := []int{}
	heap := TDAHeap.CrearHeapArr(slice, cmpEnteros)

	require.EqualValues(t, len(slice), heap.Cantidad())
	require.True(t, heap.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

	heap.Encolar(10)
	heap.Encolar(7)
	heap.Encolar(14)
	require.EqualValues(t, 14, heap.VerMax())
}

func TestHeapsortArregloVacio(t *testing.T) {
	slice := []int{}
	TDAHeap.HeapSort(slice, cmpEnteros)
	require.Equal(t, []int{}, slice, "No debe romperse por no tener elementos")
}

func TestHeapsortArregloUnElemento(t *testing.T) {
	slice := []int{10}
	TDAHeap.HeapSort(slice, cmpEnteros)
	require.Equal(t, []int{10}, slice, "El arreglo con un solo elemento debe quedar igual")
}

func TestHeapsortArreglo(t *testing.T) {
	slice := []int{4, 2, 7, 10, 4, 3, 8, 5}
	TDAHeap.HeapSort(slice, cmpEnteros)
	require.Equal(t, []int{2, 3, 4, 4, 5, 7, 8, 10}, slice, "Se ordena correctamente un arreglo")
}

func TestHeapsortArregloOrdenado(t *testing.T) {
	slice := []int{2, 3, 4, 4, 5, 7, 8, 10}
	TDAHeap.HeapSort(slice, cmpEnteros)
	require.Equal(t, []int{2, 3, 4, 4, 5, 7, 8, 10}, slice, "Un arreglo ya ordenado no cambia su orden")
}
