package lista_test

import (
	"github.com/stretchr/testify/require"
	"tdas/lista"
	"testing"
)

const (
	iteracionesAumentarVolumen  = 10000
	iteracionesDisminuirVolumen = 9999
	pruebasVolumenChico         = 5
)

func TestListaVacia(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()
	iteradorDeLista := listaEnlazada.Iterador()

	require.True(t, listaEnlazada.EstaVacia())
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestDeVolumenInsertandoPrimero(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()

	// Pruebo si puedo encolar tantos elementos como quiera.
	for i := range iteracionesAumentarVolumen {
		listaEnlazada.InsertarPrimero(i)

	}

	iteradorDeLista := listaEnlazada.Iterador()
	require.EqualValues(t, 9999, listaEnlazada.VerPrimero(), "Al aumentar el volumen el primero no coincidio el valor.")
	require.EqualValues(t, 0, listaEnlazada.VerUltimo(), "Al aumentar el volumen el ultimo no coincidio el valor.")
	require.False(t, listaEnlazada.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desencolar tantos elementos como pueda.
	for i := iteracionesDisminuirVolumen; i >= 0; i-- {
		listaEnlazada.BorrarPrimero()

	}

	iteradorDeLista = listaEnlazada.Iterador()
	require.True(t, listaEnlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestDeVolumenInsertandoUltimo(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()

	// Pruebo si puedo encolar tantos elementos como quiera.
	for i := range iteracionesAumentarVolumen {
		listaEnlazada.InsertarUltimo(i)

	}

	iteradorDeLista := listaEnlazada.Iterador()
	require.EqualValues(t, 0, listaEnlazada.VerPrimero(), "Al aumentar el volumen el primero no coincidio el valor.")
	require.EqualValues(t, 9999, listaEnlazada.VerUltimo(), "Al aumentar el volumen el ultimo no coincidio el valor.")
	require.False(t, listaEnlazada.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desencolar tantos elementos como pueda.
	for i := iteracionesDisminuirVolumen; i >= 0; i-- {
		listaEnlazada.BorrarPrimero()

	}

	iteradorDeLista = listaEnlazada.Iterador()
	require.True(t, listaEnlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorExternoAlInicio(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()

	// Pruebas insertando al inicio.
	for i := 0; i < pruebasVolumenChico; i++ {
		listaEnlazada.InsertarPrimero(i)

	}

	iteradorDeLista := listaEnlazada.Iterador()
	iteradorDeLista.Insertar(20)
	require.EqualValues(t, listaEnlazada.VerPrimero(), iteradorDeLista.VerActual(), "No coinciden los valores del primer elemento entre la lista y el iterador.")

	// Borrar el primer elemento y verificar que sea remplazado.
	iteradorDeLista.Borrar()
	require.EqualValues(t, listaEnlazada.VerPrimero(), iteradorDeLista.VerActual(), "No coinciden los valores del primer elemento entre la lista y el iterador.")

	// Limipar lista con iterador externo.
	for i := 0; i < pruebasVolumenChico; i++ {
		iteradorDeLista.Borrar()

	}

	// Pruebo que se comporte como una lista vacia.
	iteradorDeLista = listaEnlazada.Iterador()
	require.True(t, listaEnlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorExternoCentral(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()

	// Pruebas insertando en el centro. Deberia quedar algo asi: |> Primero | -> 1 -> 1 -> 5 -> 1 -> 1 | Fin |>.
	listaEnlazada.InsertarPrimero(1)
	listaEnlazada.InsertarPrimero(1)
	listaEnlazada.InsertarPrimero(1)
	listaEnlazada.InsertarPrimero(1)
	iteradorDeLista := listaEnlazada.Iterador()
	iteradorDeLista.Siguiente()
	iteradorDeLista.Siguiente()
	iteradorDeLista.Insertar(5)

	// Elimino el elemento del medio y recorro la lista verificando que no se encuentre.
	require.EqualValues(t, 5, iteradorDeLista.VerActual(), "No devuelve el valor debido.")
	iteradorDeLista.Borrar()

	iteradorDeLista = listaEnlazada.Iterador()

	for iteradorDeLista.HaySiguiente() {
		require.EqualValues(t, 1, iteradorDeLista.VerActual(), "No devuelve el valor debido.")
		iteradorDeLista.Siguiente()

	}

	iteradorDeLista = listaEnlazada.Iterador()

	// Limipar lista con iterador externo.
	for i := 0; i < (pruebasVolumenChico - 1); i++ {
		iteradorDeLista.Borrar()

	}

	// Pruebo que se comporte como una lista vacia.
	require.True(t, listaEnlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorExternoAlFinal(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()

	// Pruebas insertando al final.
	for i := 0; i < pruebasVolumenChico; i++ {
		listaEnlazada.InsertarUltimo(1)

	}

	iteradorDeLista := listaEnlazada.Iterador()
	for iteradorDeLista.HaySiguiente() {
		iteradorDeLista.Siguiente()
	}

	// Verifico que al insertar al final tambien se modifique en la lista.
	iteradorDeLista.Insertar(28)
	require.EqualValues(t, 28, iteradorDeLista.VerActual(), "No devuelve el valor esperado al insertar al final.")
	iteradorDeLista.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al VerActual de una lista vacia.")
	require.EqualValues(t, 28, listaEnlazada.VerUltimo(), "No devuelve el valor esperado al insertar al final.")

	// Limpiar la lista con iterador externo.
	iteradorDeLista = listaEnlazada.Iterador()

	for iteradorDeLista.HaySiguiente() {
		iteradorDeLista.Borrar()
	}

	// Pruebo que se comporte como una lista vacia.
	require.True(t, listaEnlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorInterno(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[int]()
	var contador int
	var arregloAux []int

	// Deberia quedar algo asi: |> Principio | 0 -> 1 -> 2 -> 3 -> 4 -> | fin |>
	for i := 0; i < pruebasVolumenChico; i++ {
		listaEnlazada.InsertarUltimo(i)

	}

	// Chequeo que la condición de corte externa se efectue.
	listaEnlazada.Iterar(func(actual int) bool {
		if contador == 3 {
			return false
		}
		actual = actual * 2
		arregloAux = append(arregloAux, actual)
		contador++
		return true
	})

	esperado := []int{0, 2, 4}
	require.EqualValues(t, esperado, arregloAux, "No devuelve el resultado esperado cuando se establece una condición de corte externa.")

	// Chequeo que se corte la iteración hasta el final.
	arregloAux = []int{}

	listaEnlazada.Iterar(func(actual int) bool {
		actual = actual * 2
		arregloAux = append(arregloAux, actual)
		return true
	})

	esperado = []int{0, 2, 4, 6, 8}
	require.EqualValues(t, esperado, arregloAux, "No devuelve el valor esperado al iterar hasta el final.")

	// Prueba con corte instantaneo.
	arregloAux = []int{}

	listaEnlazada.Iterar(func(actual int) bool {
		return false
	})

	esperado = []int{}
	require.EqualValues(t, esperado, arregloAux, "No devuelve el valor esperado al iterar hasta el final.")

}

func TestListaDeStrings(t *testing.T) {
	listaEnlazada := lista.CrearListaEnlazada[string]()
	iteradorDeLista := listaEnlazada.Iterador()

	var (
		a = "hola"
		b = "como"
		c = "estas"
		d = "?"
	)

	// Pruebo Insertar a lo ultimo cadenas y verificar si se borran correctamente.
	listaEnlazada.InsertarUltimo(a)
	listaEnlazada.InsertarUltimo(b)
	listaEnlazada.InsertarUltimo(c)
	listaEnlazada.InsertarUltimo(d)
	require.False(t, listaEnlazada.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo Borrar todas las cadenas y verifico los primeros.
	require.EqualValues(t, a, listaEnlazada.VerPrimero())
	require.EqualValues(t, a, listaEnlazada.BorrarPrimero())

	require.EqualValues(t, b, listaEnlazada.VerPrimero())
	require.EqualValues(t, b, listaEnlazada.BorrarPrimero())

	require.EqualValues(t, c, listaEnlazada.VerPrimero())
	require.EqualValues(t, c, listaEnlazada.BorrarPrimero())

	require.EqualValues(t, d, listaEnlazada.VerPrimero())
	require.EqualValues(t, d, listaEnlazada.BorrarPrimero())

	// Pruebo que al Borrar todas las cadenas se comporte como una lista vacia.
	require.True(t, listaEnlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, listaEnlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}
