package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"
)

var TAMANIOSVOLUMEN = []int{12500, 25000, 50000, 100000, 200000}

func funcionCmpInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}

	return 0
}

func funcionCmpString(a, b string) int {
	return strings.Compare(a, b)
}

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })
}

func TestABBUnElemento(t *testing.T) {
	t.Log("Comprueba que diccionario con un elemento tiene esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	abb.Guardar(1, 99)

	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, 99, abb.Obtener(1))
	require.True(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(2) })
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

	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)
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

func TestABBRemplazarDatos(t *testing.T) {
	t.Log("Guarda una clave, y luego vuelve a guardar esta misma, buscando que el dato se haya reemplazado")
	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	abb.Guardar(1, 101)
	abb.Guardar(1, 102)
	require.EqualValues(t, 102, abb.Obtener(1))

	abb.Guardar(1, 103)
	require.EqualValues(t, 103, abb.Obtener(1))

	require.EqualValues(t, 1, abb.Cantidad())

}

func TestABBReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)
	abb.Guardar(100, 100)
	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(10, 10)
	abb.Guardar(75, 75)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(150, 150)
	abb.Guardar(125, 125)
	abb.Guardar(110, 110)
	abb.Guardar(200, 200)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)

	abb.Guardar(100, 100*2)
	abb.Guardar(50, 50*2)
	abb.Guardar(25, 25*2)
	abb.Guardar(10, 10*2)
	abb.Guardar(75, 75*2)
	abb.Guardar(60, 60*2)
	abb.Guardar(80, 80*2)
	abb.Guardar(150, 150*2)
	abb.Guardar(125, 125*2)
	abb.Guardar(110, 110*2)
	abb.Guardar(200, 200*2)
	abb.Guardar(190, 190*2)
	abb.Guardar(300, 300*2)

	claves := []int{100, 50, 25, 10, 75, 60, 80, 150, 125, 110, 200, 190, 300}
	for _, valor := range claves {
		require.EqualValues(t, valor*2, abb.Obtener(valor))
	}

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
	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
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

func TestABBReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)
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

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb := TDADiccionario.CrearABB[int, string](funcionCmpInt)
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
	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDADiccionario.CrearABB[string, *int](funcionCmpString)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar")

	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)
	for i := 0; i < 1000; i++ {
		abb.Guardar(i, i)
		require.True(t, abb.Pertenece(i))
		abb.Borrar(i)
		require.False(t, abb.Pertenece(i))
	}
}

func buscarEnABB(clave string, claves []string) int {
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
	abb := TDADiccionario.CrearABB[string, *int](funcionCmpString)
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
	require.NotEqualValues(t, -1, buscarEnABB(cs[0], claves))
	require.NotEqualValues(t, -1, buscarEnABB(cs[1], claves))
	require.NotEqualValues(t, -1, buscarEnABB(cs[2], claves))
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

	abb := TDADiccionario.CrearABB[string, int](funcionCmpString)
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

	abb := TDADiccionario.CrearABB[string, int](funcionCmpString)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestABBIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	abb := TDADiccionario.CrearABB[string, int](funcionCmpString)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDeRangoIterdaroInterno(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno con rango definido, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	abb.Guardar(100, 100)
	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(10, 10)
	abb.Guardar(75, 75)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(150, 150)
	abb.Guardar(125, 125)
	abb.Guardar(110, 110)
	abb.Guardar(200, 200)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)

	desde, hasta := 50, 200

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	abb.IterarRango(&desde, &hasta, func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c == 75 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestABBDiccionarioIterar(t *testing.T) {
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
	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarEnABB(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarEnABB(segundo, claves))
	require.EqualValues(t, valores[buscarEnABB(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarEnABB(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)
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
	require.NotEqualValues(t, -1, buscarEnABB(primero, claves))
	require.NotEqualValues(t, -1, buscarEnABB(segundo, claves))
	require.NotEqualValues(t, -1, buscarEnABB(tercero, claves))
}

func TestABBPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del diccionario abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	abb := TDADiccionario.CrearABB[string, string](funcionCmpString)
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

func TestVolumenIteradorCorteABB(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	/* Inserta 'n' parejas en el diccionario */
	abb.Guardar(100, 100)
	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(10, 10)
	abb.Guardar(75, 75)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(150, 150)
	abb.Guardar(125, 125)
	abb.Guardar(110, 110)
	abb.Guardar(200, 200)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	abb.Iterar(func(c int, v int) bool {
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
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestDeRangoIteradorExterno(t *testing.T) {
	t.Log("Prueba de volumen de iterador Externo con rango definido, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	abb.Guardar(100, 100)
	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(10, 10)
	abb.Guardar(75, 75)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(150, 150)
	abb.Guardar(125, 125)
	abb.Guardar(110, 110)
	abb.Guardar(200, 200)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)

	desde, hasta := 60, 120
	esperados, indice := []int{60, 75, 80, 100, 110}, 0

	for iter := abb.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		fmt.Println(clave)
		fmt.Println(esperados[indice])
		require.EqualValues(t, esperados[indice], clave, "No se devuelve el valor esperado.")
		indice++
	}

}

// Este test fue creado en base a la información del error que nos dio rosita.
func TestIterarRangoCasoError(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](funcionCmpInt)

	abb.Guardar(1, 1)
	abb.Guardar(2, 2)
	abb.Guardar(3, 3)
	abb.Guardar(4, 4)
	abb.Guardar(5, 5)
	abb.Guardar(6, 6)
	abb.Guardar(7, 7)

	desde := 2
	hasta := 5
	esperado := 2

	abb.IterarRango(&desde, &hasta, func(clave int, _ int) bool {
		require.EqualValues(t, esperado, clave, "no")
		esperado++
		return true
	})
}
