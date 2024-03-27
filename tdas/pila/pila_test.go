package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "No emite error al Desapilar pila vacia.")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "No emite error al VerTope de pila vacia.")
}

func TestDeVolumenDePila(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	// Pruebo que puedo apilar tantos elementos como quiera.
	for i := range 1000 {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope(), "Al aumentar el volumen no coincidio el tope.")
	}

	require.False(t, pila.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo que puedo desapilar tantos elementos como pueda.
	for i := 998; i >= 0; i-- {
		pila.Desapilar()
		require.EqualValues(t, i, pila.VerTope(), "Al disminuir el volumen no coincidio el tope.")
	}

	// Verifico las invariantes.
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Al reducir el volumen no ejecuta el error.")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Al reducir el volumen no ejecuta el error.")
	require.True(t, pila.EstaVacia(), "No se comporta como recien creada.")
}

func TestPilaDeStrings(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[string]()

	var (
		a = "hola"
		b = "como"
		c = "estas"
		d = "?"
	)

	// Pruebo apilar cadenas y verificar si se apilan correctamente.
	pila.Apilar(a)
	require.EqualValues(t, "hola", pila.VerTope())
	pila.Apilar(b)
	require.EqualValues(t, "como", pila.VerTope())
	pila.Apilar(c)
	require.EqualValues(t, "estas", pila.VerTope())
	pila.Apilar(d)
	require.EqualValues(t, "?", pila.VerTope())
	require.False(t, pila.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo desapilar todas las cadenas y verificar los topes, tambien chequeo si al desapilar todo se comporta como una pila vacia.
	pila.Desapilar()
	require.EqualValues(t, "estas", pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, "como", pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, "hola", pila.VerTope())

	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}
