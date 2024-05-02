package lista_test

import (
	"github.com/stretchr/testify/require"
	lista "tdas/lista"
	"testing"
)

const (
	_ITERACIONES_AUMENTAR_VOLUMEN  = 10000
	_ITERACIONES_DISMINUIR_VOLUMEN = 9999
	_PRUEBAS_VOLUMEN_CHICO         = 5
)

func TestListaVacia(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()
	iteradorDeLista := lista_Enlazada.Iterador()

	require.True(t, lista_Enlazada.EstaVacia())
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestDeVolumenInsertandoPrimero(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()

	// Pruebo si puedo encolar tantos elementos como quiera.
	for i := range _ITERACIONES_AUMENTAR_VOLUMEN {
		lista_Enlazada.InsertarPrimero(i)

	}

	iteradorDeLista := lista_Enlazada.Iterador()
	require.EqualValues(t, 9999, lista_Enlazada.VerPrimero(), "Al aumentar el volumen el primero no coincidio el valor.")
	require.EqualValues(t, 0, lista_Enlazada.VerUltimo(), "Al aumentar el volumen el ultimo no coincidio el valor.")
	require.False(t, lista_Enlazada.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desencolar tantos elementos como pueda.
	for i := _ITERACIONES_DISMINUIR_VOLUMEN; i >= 0; i-- {
		lista_Enlazada.BorrarPrimero()

	}

	iteradorDeLista = lista_Enlazada.Iterador()
	require.True(t, lista_Enlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestDeVolumenInsertandoUltimo(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()

	// Pruebo si puedo encolar tantos elementos como quiera.
	for i := range _ITERACIONES_AUMENTAR_VOLUMEN {
		lista_Enlazada.InsertarUltimo(i)

	}

	iteradorDeLista := lista_Enlazada.Iterador()
	require.EqualValues(t, 0, lista_Enlazada.VerPrimero(), "Al aumentar el volumen el primero no coincidio el valor.")
	require.EqualValues(t, 9999, lista_Enlazada.VerUltimo(), "Al aumentar el volumen el ultimo no coincidio el valor.")
	require.False(t, lista_Enlazada.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desencolar tantos elementos como pueda.
	for i := _ITERACIONES_DISMINUIR_VOLUMEN; i >= 0; i-- {
		lista_Enlazada.BorrarPrimero()

	}

	iteradorDeLista = lista_Enlazada.Iterador()
	require.True(t, lista_Enlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorExternoAlInicio(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()

	// Pruebas insertando al inicio.
	for i := 0; i < _PRUEBAS_VOLUMEN_CHICO; i++ {
		lista_Enlazada.InsertarPrimero(i)

	}

	iteradorDeLista := lista_Enlazada.Iterador()
	iteradorDeLista.Insertar(20)
	require.EqualValues(t, lista_Enlazada.VerPrimero(), iteradorDeLista.VerActual(), "No coinciden los valores del primer elemento entre la lista y el iterador.")

	// Borrar el primer elemento y verificar que sea remplazado.
	iteradorDeLista.Borrar()
	require.EqualValues(t, lista_Enlazada.VerPrimero(), iteradorDeLista.VerActual(), "No coinciden los valores del primer elemento entre la lista y el iterador.")

	// Limipar lista con iterador externo.
	for i := 0; i < _PRUEBAS_VOLUMEN_CHICO; i++ {
		iteradorDeLista.Borrar()

	}

	// Pruebo que se comporte como una lista vacia.
	iteradorDeLista = lista_Enlazada.Iterador()
	require.True(t, lista_Enlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorExternoCentral(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()

	// Pruebas insertando en el centro. Deberia quedar algo asi: |> Primero | -> 1 -> 1 -> 5 -> 1 -> 1 | Fin |>.
	lista_Enlazada.InsertarPrimero(1)
	lista_Enlazada.InsertarPrimero(1)
	lista_Enlazada.InsertarPrimero(1)
	lista_Enlazada.InsertarPrimero(1)
	iteradorDeLista := lista_Enlazada.Iterador()
	iteradorDeLista.Siguiente()
	iteradorDeLista.Siguiente()
	iteradorDeLista.Insertar(5)

	// Elimino el elemento del medio y recorro la lista verificando que no se encuentre.
	require.EqualValues(t, 5, iteradorDeLista.VerActual(), "No devuelve el valor debido.")
	iteradorDeLista.Borrar()

	iteradorDeLista = lista_Enlazada.Iterador()

	for iteradorDeLista.HaySiguiente() {
		require.EqualValues(t, 1, iteradorDeLista.VerActual(), "No devuelve el valor debido.")
		iteradorDeLista.Siguiente()

	}

	iteradorDeLista = lista_Enlazada.Iterador()

	// Limipar lista con iterador externo.
	for i := 0; i < (_PRUEBAS_VOLUMEN_CHICO - 1); i++ {
		iteradorDeLista.Borrar()

	}

	// Pruebo que se comporte como una lista vacia.
	require.True(t, lista_Enlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorExternoAlFinal(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()

	// Pruebas insertando al final.
	for i := 0; i < _PRUEBAS_VOLUMEN_CHICO; i++ {
		lista_Enlazada.InsertarUltimo(1)

	}

	iteradorDeLista := lista_Enlazada.Iterador()
	for iteradorDeLista.HaySiguiente() {
		iteradorDeLista.Siguiente()

	}

	// Verifico que al insertar al final tambien se modifique en la lista.
	iteradorDeLista.Insertar(28)
	iteradorDeLista.Siguiente()
	require.EqualValues(t, lista_Enlazada.VerUltimo(), iteradorDeLista.VerActual(), "No coinceden en el valor al posicionarse en el final la lista y el iterador.")

	// Verifico que al eliminar el ultimo tambien se modifique en la lista.
	iteradorDeLista.Borrar()

	iteradorDeLista = lista_Enlazada.Iterador()
	for iteradorDeLista.HaySiguiente() {
		iteradorDeLista.Siguiente()

	}

	require.EqualValues(t, lista_Enlazada.VerUltimo(), iteradorDeLista.VerActual(), "No coinceden en el valor al posicionarse en el final la lista y el iterador.")

	// Limpiar la lista con iterador externo.
	iteradorDeLista = lista_Enlazada.Iterador()
	for i := 0; i < _PRUEBAS_VOLUMEN_CHICO; i++ {
		iteradorDeLista.Borrar()

	}

	// Pruebo que se comporte como una lista vacia.
	require.True(t, lista_Enlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}

func TestIteradorInterno(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[int]()
	var contador int
	var arregloAux []int

	// Deberia quedar algo asi: |> Principio | 0 -> 1 -> 2 -> 3 -> 4 -> | fin |>
	for i := 0; i < _PRUEBAS_VOLUMEN_CHICO; i++ {
		lista_Enlazada.InsertarUltimo(i)

	}

	// Chequeo que la condición de corte externa se efectue.
	lista_Enlazada.Iterar(func(actual int) bool {
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

	lista_Enlazada.Iterar(func(actual int) bool {
		actual = actual * 2
		arregloAux = append(arregloAux, actual)
		return true
	})

	esperado = []int{0, 2, 4, 6, 8}
	require.EqualValues(t, esperado, arregloAux, "No devuelve el valor esperado al iterar hasta el final.")

	// Prueba con corte instantaneo.
	arregloAux = []int{}

	lista_Enlazada.Iterar(func(actual int) bool {
		return false
	})

	esperado = []int{}
	require.EqualValues(t, esperado, arregloAux, "No devuelve el valor esperado al iterar hasta el final.")

}

func TestListaDeStrings(t *testing.T) {
	lista_Enlazada := lista.CrearListaEnlazada[string]()
	iteradorDeLista := lista_Enlazada.Iterador()

	var (
		a = "hola"
		b = "como"
		c = "estas"
		d = "?"
	)

	// Pruebo Insertar a lo ultimo cadenas y verificar si se borran correctamente.
	lista_Enlazada.InsertarUltimo(a)
	lista_Enlazada.InsertarUltimo(b)
	lista_Enlazada.InsertarUltimo(c)
	lista_Enlazada.InsertarUltimo(d)
	require.False(t, lista_Enlazada.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo Borrar todas las cadenas y verifico los primeros.
	require.EqualValues(t, a, lista_Enlazada.VerPrimero())
	require.EqualValues(t, a, lista_Enlazada.BorrarPrimero())

	require.EqualValues(t, b, lista_Enlazada.VerPrimero())
	require.EqualValues(t, b, lista_Enlazada.BorrarPrimero())

	require.EqualValues(t, c, lista_Enlazada.VerPrimero())
	require.EqualValues(t, c, lista_Enlazada.BorrarPrimero())

	require.EqualValues(t, d, lista_Enlazada.VerPrimero())
	require.EqualValues(t, d, lista_Enlazada.BorrarPrimero())

	// Pruebo que al Borrar todas las cadenas se comporte como una lista vacia.
	require.True(t, lista_Enlazada.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.EqualValues(t, 0, lista_Enlazada.Largo(), "No devuelve el largo esperado.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.BorrarPrimero() }, "No emite error al Borrar el primer elemento.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorDeLista.Borrar() }, "No emite error al Borrar algun elemento.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerPrimero() }, "No emite error al VerPrimero de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista_Enlazada.VerUltimo() }, "No emite error al VerUltimo de una lista vacia.")
	require.PanicsWithValue(t, "La lista esta vacia", func() { iteradorDeLista.VerActual() }, "No emite error al VerActual de una lista vacia.")
}
