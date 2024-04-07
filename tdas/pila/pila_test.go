package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

const _ITERACIONESAUMENTARVOLUMEN = 10000
const _ITERACIONESDISMINUIRVOLUMEN = 9999

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "No emite error al Desapilar pila vacia.")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "No emite error al VerTope de pila vacia.")
}

func TestDeVolumenDePila(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	// Pruebo si puedo apilar tantos elementos como quiera.
	for i := range _ITERACIONESAUMENTARVOLUMEN {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope(), "Al aumentar el volumen no coincidio el tope.")
	}

	require.False(t, pila.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desapilar tantos elementos como pueda.
	for i := _ITERACIONESDISMINUIRVOLUMEN; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope(), "Al disminuir el volumen no coincidio el tope.")
		pila.Desapilar()
	}

	// Verifico las invariantes.
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
	require.EqualValues(t, "?", pila.Desapilar())
	require.EqualValues(t, "estas", pila.VerTope())
	require.EqualValues(t, "estas", pila.Desapilar())
	require.EqualValues(t, "como", pila.VerTope())
	require.EqualValues(t, "como", pila.Desapilar())
	require.EqualValues(t, "hola", pila.VerTope())

	require.EqualValues(t, "hola", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

func TestPilaDeBool(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[bool]()

	var (
		a = true
		b = false
		c = true
		d = false
	)

	// Pruebo apilar cadenas y verificar si se apilan correctamente.
	pila.Apilar(a)
	require.EqualValues(t, true, pila.VerTope())
	pila.Apilar(b)
	require.EqualValues(t, false, pila.VerTope())
	pila.Apilar(c)
	require.EqualValues(t, true, pila.VerTope())
	pila.Apilar(d)
	require.EqualValues(t, false, pila.VerTope())
	require.False(t, pila.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo desapilar todas las cadenas y verificar los topes, tambien chequeo si al desapilar todo se comporta como una pila vacia.
	require.EqualValues(t, false, pila.Desapilar())
	require.EqualValues(t, true, pila.VerTope())
	require.EqualValues(t, true, pila.Desapilar())
	require.EqualValues(t, false, pila.VerTope())
	require.EqualValues(t, false, pila.Desapilar())
	require.EqualValues(t, true, pila.VerTope())
	require.EqualValues(t, true, pila.Desapilar())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

func TestPilaDePilas(t *testing.T) {

	pilaDePilas := TDAPila.CrearPilaDinamica[TDAPila.Pila[int]]()
	pila1 := TDAPila.CrearPilaDinamica[int]()
	pila2 := TDAPila.CrearPilaDinamica[int]()
	pila3 := TDAPila.CrearPilaDinamica[int]()

	// Relleno las diferentes pilas con diferentes numeros para diferenciarlos
	for i := range 6 {
		pila1.Apilar(i)
		pila2.Apilar(i + 5)
		pila3.Apilar(i + 10)
	}

	// Pruebo apilar las pilas.
	pilaDePilas.Apilar(pila1)
	require.EqualValues(t, pila1, pilaDePilas.VerTope())
	pilaDePilas.Apilar(pila2)
	require.EqualValues(t, pila2, pilaDePilas.VerTope())
	pilaDePilas.Apilar(pila3)
	require.EqualValues(t, pila3, pilaDePilas.VerTope())

	// Pruebo a desapilar e investigar sus elementos posteriormente.
	elemento := pilaDePilas.Desapilar()
	require.EqualValues(t, 15, elemento.Desapilar())
	require.EqualValues(t, 14, elemento.Desapilar())
	require.EqualValues(t, 13, elemento.Desapilar())
	require.EqualValues(t, 12, elemento.Desapilar())
	require.EqualValues(t, 11, elemento.Desapilar())
	elemento = pilaDePilas.Desapilar()
	require.EqualValues(t, 10, elemento.Desapilar())
	require.EqualValues(t, 9, elemento.Desapilar())
	require.EqualValues(t, 8, elemento.Desapilar())
	require.EqualValues(t, 7, elemento.Desapilar())
	require.EqualValues(t, 6, elemento.Desapilar())
	elemento = pilaDePilas.Desapilar()
	require.EqualValues(t, 5, elemento.Desapilar())
	require.EqualValues(t, 4, elemento.Desapilar())
	require.EqualValues(t, 3, elemento.Desapilar())
	require.EqualValues(t, 2, elemento.Desapilar())
	require.EqualValues(t, 1, elemento.Desapilar())

	// Verifico que funcione como una pila vacia.
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaDePilas.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaDePilas.VerTope() })
	require.True(t, pilaDePilas.EstaVacia())

}
