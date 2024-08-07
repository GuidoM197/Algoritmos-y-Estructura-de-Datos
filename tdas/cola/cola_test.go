package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

const _ITERACIONES_AUMENTAR_VOLUMEN = 10000
const _ITERACIONES_DISMINUIR_VOLUMEN = 9999

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "No emite error al Desencolar cola vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "No emite error al VerPrimero de cola vacia.")
}

func TestDeVolumenDeCola(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	// Pruebo si puedo encolar tantos elementos como quiera.
	for i := range _ITERACIONES_AUMENTAR_VOLUMEN {
		cola.Encolar(i)
	}
	require.EqualValues(t, 0, cola.VerPrimero(), "Al aumentar el volumen no coincidio el primero.")

	require.False(t, cola.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desencolar tantos elementos como pueda.
	for i := _ITERACIONES_DISMINUIR_VOLUMEN; i >= 0; i-- {
		cola.Desencolar()
	}

	require.True(t, cola.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "No emite error al Desencolar cola vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "No emite error al VerPrimero de cola vacia.")
}

func TestColaDeStrings(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[string]()

	var (
		a = "hola"
		b = "como"
		c = "estas"
		d = "?"
	)

	// Pruebo apilar cadenas y verificar si se encolar correctamente.
	cola.Encolar(a)
	cola.Encolar(b)
	cola.Encolar(c)
	cola.Encolar(d)
	require.False(t, cola.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo desencolar todas las cadenas y verifico los primeros.
	require.EqualValues(t, "hola", cola.VerPrimero())
	require.EqualValues(t, "hola", cola.Desencolar())

	require.EqualValues(t, "como", cola.VerPrimero())
	require.EqualValues(t, "como", cola.Desencolar())

	require.EqualValues(t, "estas", cola.VerPrimero())
	require.EqualValues(t, "estas", cola.Desencolar())

	require.EqualValues(t, "?", cola.VerPrimero())
	require.EqualValues(t, "?", cola.Desencolar())

	// Pruebo que al desencolar todo se comporte como una cola vacia.
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestColaaDeBool(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[bool]()

	var (
		a = true
		b = false
		c = true
		d = false
	)

	// Pruebo apilar cadenas y verificar si se apilan correctamente.
	cola.Encolar(a)
	cola.Encolar(b)
	cola.Encolar(c)
	cola.Encolar(d)
	require.False(t, cola.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo desencolar todas las cadenas y verificar los primeros.
	require.EqualValues(t, true, cola.VerPrimero())
	require.EqualValues(t, true, cola.Desencolar())

	require.EqualValues(t, false, cola.VerPrimero())
	require.EqualValues(t, false, cola.Desencolar())

	require.EqualValues(t, true, cola.VerPrimero())
	require.EqualValues(t, true, cola.Desencolar())

	require.EqualValues(t, false, cola.VerPrimero())
	require.EqualValues(t, false, cola.Desencolar())

	// Pruebo que al desencolar se comporte como una cola vacia.
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}
