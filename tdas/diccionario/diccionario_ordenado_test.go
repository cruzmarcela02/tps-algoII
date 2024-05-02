package diccionario_test

import (
	"fmt"
	"strings"
	TDAAbb "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{1000, 5000, 12000}

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que diccionario vacio no tiene claves")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(10) })
}

func cmpEnteros(clave1, clave2 int) int {
	return clave1 - clave2
}

func cmpStrings(clave1, clave2 string) int {
	return strings.Compare(clave1, clave2)
}

func TestABBClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un diccionario vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDAAbb.CrearABB[int, string](cmpEnteros)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestABBUnElement(t *testing.T) {
	t.Log("Comprueba que diccionario con un elemento tiene esa Clave, unicamente")
	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestABBReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	abb := TDAAbb.CrearABB[int, int](cmpEnteros)

	for i := 0; i < 500; i++ {
		abb.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestABBBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDAAbb.CrearABB[string, string](cmpStrings)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })
}

func TestABBReutilizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	clave := "hola"
	abb.Guardar(clave, "mundo!")
	abb.Borrar(clave)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(clave))
	abb.Guardar(clave, "mundooo!")
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, "mundooo!", abb.Obtener(clave))
}

func TestABBBorrarDosHijosRaiz(t *testing.T) {
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"
	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	require.EqualValues(t, 6, abb.Cantidad())
	require.EqualValues(t, 7, abb.Borrar(clave0))
	require.EqualValues(t, 5, abb.Cantidad())
}

func TestABBBorrarDosHijos(t *testing.T) {
	clave0 := 10
	clave1 := 5
	clave2 := 8
	clave3 := 15
	clave4 := 2
	clave5 := 11
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(clave0, 10)
	abb.Guardar(clave1, 5)
	abb.Guardar(clave2, 8)
	abb.Guardar(clave3, 15)
	abb.Guardar(clave4, 2)
	abb.Guardar(clave5, 11)

	require.EqualValues(t, 5, abb.Borrar(clave1))
	require.EqualValues(t, 5, abb.Cantidad())
}

func TestABBBorrarDosHijosMuchosElementos(t *testing.T) {
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(10, 10)
	abb.Guardar(5, 5)
	abb.Guardar(8, 8)
	abb.Guardar(15, 15)
	abb.Guardar(2, 2)
	abb.Guardar(11, 11)
	abb.Guardar(30, 30)
	abb.Guardar(20, 20)
	abb.Guardar(60, 60)
	abb.Guardar(50, 50)
	abb.Guardar(70, 70)

	require.EqualValues(t, 60, abb.Borrar(60))
	require.EqualValues(t, 10, abb.Cantidad())
}

func TestABBBorrarUnHijo(t *testing.T) {
	clave0 := 10
	clave1 := 5
	clave2 := 8
	clave3 := 15
	clave4 := 2
	clave5 := 11
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(clave0, 10)
	abb.Guardar(clave1, 5)
	abb.Guardar(clave2, 8)
	abb.Guardar(clave3, 15)
	abb.Guardar(clave4, 2)
	abb.Guardar(clave5, 11)

	require.EqualValues(t, 2, abb.Borrar(clave4))
	require.EqualValues(t, 5, abb.Cantidad())
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb := TDAAbb.CrearABB[int, string](cmpEnteros)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDAAbb.CrearABB[string, *int](cmpStrings)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		abb.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, abb.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = abb.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestABBIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDAAbb.CrearABB[string, *int](cmpStrings)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestABBIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	require.EqualValues(t, 7, abb.Borrar(clave0))

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestABBIteradorInternoPorRango(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno por rango desde y hasta")
	clave0 := 10
	clave1 := 5
	clave2 := 8
	clave3 := 15
	clave4 := 2
	clave5 := 11
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(clave0, 10)
	abb.Guardar(clave1, 5)
	abb.Guardar(clave2, 8)
	abb.Guardar(clave3, 15)
	abb.Guardar(clave4, 2)
	abb.Guardar(clave5, 11)

	suma := 0
	ptrSuma := &suma
	desde := 5
	hasta := 10
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return true
	})

	require.EqualValues(t, 23, suma)
}

func TestABBIteradorInternoFueraDeRango(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno por rango desde y hasta")
	clave0 := 10
	clave1 := 5
	clave2 := 8
	clave3 := 15
	clave4 := 2
	clave5 := 11
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(clave0, 10)
	abb.Guardar(clave1, 5)
	abb.Guardar(clave2, 8)
	abb.Guardar(clave3, 15)
	abb.Guardar(clave4, 2)
	abb.Guardar(clave5, 11)

	suma := 0
	ptrSuma := &suma
	desde := 150
	hasta := 4320
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return true
	})

	require.EqualValues(t, 0, suma)
}

func TestABBIteradorCondicionCorte(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	suma := 0
	ptrSuma := &suma
	abb.Iterar(func(_ string, dato int) bool {
		*ptrSuma += dato
		return suma < 15
	})

	require.EqualValues(t, 15, suma)
}

func TestABBIteradorCondicionCorteConRango(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	suma := 0
	ptrSuma := &suma

	abb.IterarRango(&clave5, &clave3, func(_ string, dato int) bool {
		*ptrSuma += dato
		return suma < 6
	})

	require.EqualValues(t, 7, suma)
}

func TestABBIteradorRangoGrande(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno por rango desde y hasta")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(100, 100)

	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(75, 75)
	abb.Guardar(10, 10)
	abb.Guardar(30, 30)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)

	abb.Guardar(150, 150)
	abb.Guardar(125, 125)
	abb.Guardar(200, 200)
	abb.Guardar(110, 110)
	abb.Guardar(130, 130)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)

	suma := 0
	ptrSuma := &suma
	desde := 30
	hasta := 110
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return suma < 505
	})

	require.EqualValues(t, 505, suma)
}

func TestABBIteradorRangoNilToHasta(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno por rango desde = nil y hasta " +
		"Sin embargo si la condicion de corte no se cumple, corta la iteracion sin necesidad de haber llegado al hasta")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(75, 75)
	abb.Guardar(10, 10)
	abb.Guardar(30, 30)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(150, 150)
	abb.Guardar(125, 125)
	abb.Guardar(200, 200)
	abb.Guardar(110, 110)
	abb.Guardar(130, 130)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)

	suma := 0
	ptrSuma := &suma
	hasta := 75
	abb.IterarRango(nil, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return suma < 250
	})

	require.EqualValues(t, 250, suma)
}

func TestABBIteradorCondicionDesdeToNil(t *testing.T) {
	t.Log("Itera sobre el  diccionario desde una cierta clave, hasta el final, sin empabargo si la condicion no se cumple " +
		"se corta la iteracion sin necesitdad de iterar hasta el final ")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	desde := clave5
	suma := 0
	ptrSuma := &suma
	abb.IterarRango(&desde, nil, func(_ string, dato int) bool {
		*ptrSuma += dato
		return suma < 7
	})

	require.EqualValues(t, 7, suma)
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	abb := TDAAbb.CrearABB[string, int](cmpStrings)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = abb.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = abb.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = abb.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, abb.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestABBIteradorDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	abb := TDAAbb.CrearABB[string, int](cmpStrings)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorRangos(t *testing.T) {
	t.Log("Iterar sobre diccionario desde el 40 al 60")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	c1 := 50
	c2 := 100
	c3 := 110
	abb.Guardar(c2, 1)
	abb.Guardar(c1, 0)
	abb.Guardar(c3, 1)
	desde := 40
	hasta := 60
	iter := abb.IteradorRango(&desde, &hasta)
	require.True(t, iter.HaySiguiente())

	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		require.EqualValues(t, 0, v)
		iter.Siguiente()
	}
}

func TestABBIteradorRangoMAYOR(t *testing.T) {
	t.Log("Iterar sobre diccionario desde el 105 al 110")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	c1 := 50
	c2 := 100
	c3 := 110
	abb.Guardar(c2, 0)
	abb.Guardar(c1, 0)
	abb.Guardar(c3, 1)
	desde := 105
	hasta := 110
	iter := abb.IteradorRango(&desde, &hasta)
	require.True(t, iter.HaySiguiente())

	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		require.EqualValues(t, 1, v)
		iter.Siguiente()
	}
}
func TestABBIteradorRangosMasGrande(t *testing.T) {
	t.Log("Iterar sobre diccionario desde el 60 al 120")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)
	abb.Guardar(14, 1)
	abb.Guardar(4, 1)
	abb.Guardar(3, 1)
	abb.Guardar(9, 1)
	abb.Guardar(22, 0)
	abb.Guardar(16, 1)
	abb.Guardar(24, 0)
	desde := 3
	hasta := 16
	iter := abb.IteradorRango(&desde, &hasta)
	require.True(t, iter.HaySiguiente())

	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		require.EqualValues(t, 1, v)
		iter.Siguiente()
	}
}
func TestABBDiccionarioOrdenadoIterador(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	claves := []string{"A", "B", "C"}
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	abb.Iterador()
	iter2 := abb.Iterador()
	iter2.Siguiente()
	iter3 := abb.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestABBPruebaIteradorTrasBorrados(t *testing.T) {
	t.Log("Esta prueba intenta verificar el comportamiento del iterador del Abb que ha guardado 3 veces pero las ha borrado " +
		"el iterador no deberia ejecutar sus funciones. Sin embargo al guardar nuevamente, y generar un nuevo iterador " +
		"se comprueba las claves y los datos guardados.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	abb := TDAAbb.CrearABB[string, string](cmpStrings)
	abb.Guardar(clave1, "")
	abb.Guardar(clave2, "")
	abb.Guardar(clave3, "")
	abb.Borrar(clave1)
	abb.Borrar(clave2)
	abb.Borrar(clave3)
	iter := abb.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	abb.Guardar(clave1, "A")
	iter = abb.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestABBIteradorRangosEJClase(t *testing.T) {
	t.Log("Iterar sobre diccionario desde el nil al 9")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)

	abb.Guardar(14, 0)
	abb.Guardar(4, 1)
	abb.Guardar(3, 1)
	abb.Guardar(9, 1)
	abb.Guardar(22, 0)
	abb.Guardar(16, 0)
	abb.Guardar(24, 0)

	hasta := 9
	iter := abb.IteradorRango(nil, &hasta)
	require.True(t, iter.HaySiguiente())

	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		require.EqualValues(t, 1, v)
		iter.Siguiente()
	}
}

func TestABBIteradorRangosNilToNil(t *testing.T) {
	t.Log("Iterar sobre diccionario desde el inici a fin, pues se le pasa nil por parametro")
	abb := TDAAbb.CrearABB[int, int](cmpEnteros)

	abb.Guardar(14, 1)
	abb.Guardar(4, 1)
	abb.Guardar(3, 1)
	abb.Guardar(9, 1)
	abb.Guardar(22, 1)
	abb.Guardar(16, 1)
	abb.Guardar(24, 1)
	iter := abb.IteradorRango(nil, nil)
	require.True(t, iter.HaySiguiente())

	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		require.EqualValues(t, 1, v)
		iter.Siguiente()
	}
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	abb := TDAAbb.CrearABB[string, *int](cmpStrings)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		abb.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := abb.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDAAbb.CrearABB[int, int](cmpEnteros)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}
