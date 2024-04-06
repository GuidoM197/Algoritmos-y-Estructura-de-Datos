package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	TDAPila "tdas/pila"
	"testing"
)

const IteracionesAumentarVolumen = 10000
const IteracionesDisminuirVolumen = 9999

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.Desencolar() }, "No emite error al Desencolar pila vacia.")
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.VerPrimero() }, "No emite error al VerPrimero de pila vacia.")
}

func TestDeVolumenDeCola(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	// Pruebo si puedo encolar tantos elementos como quiera.
	for i := range IteracionesAumentarVolumen {
		cola.Encolar(i)
	}
	require.EqualValues(t, 0, cola.VerPrimero(), "Al aumentar el volumen no coincidio el primero.")

	require.False(t, cola.EstaVacia(), "Deberia devolver False ya que contiene elementos.")

	// Pruebo si puedo desencolar tantos elementos como pueda.
	for i := IteracionesDisminuirVolumen; i >= 0; i-- {
		cola.Desencolar()
	}

	require.True(t, cola.EstaVacia(), "Deberia devolver True ya que NO contiene elementos.")
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
	require.EqualValues(t, "hola", cola.VerPrimero())
	require.False(t, cola.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo desencolar todas las cadenas y verificar los topes.
	require.EqualValues(t, "hola", cola.Desencolar())
	require.EqualValues(t, "como", cola.Desencolar())
	require.EqualValues(t, "estas", cola.Desencolar())

	require.EqualValues(t, "?", cola.VerPrimero())
	require.EqualValues(t, "?", cola.Desencolar())

	// Pruebo que al desencolar todo se comporte como una cola vacia.
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestPilaDeBool(t *testing.T) {

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
	require.EqualValues(t, true, cola.VerPrimero())
	require.False(t, cola.EstaVacia(), "Si contiene elementos debe devolver False.")

	// Pruebo desencolar todas las cadenas y verificar los topes.
	require.EqualValues(t, true, cola.Desencolar())
	require.EqualValues(t, false, cola.Desencolar())
	require.EqualValues(t, true, cola.Desencolar())
	require.EqualValues(t, false, cola.Desencolar())

	// Pruebo que al desencolar se comporte como una cola vacia.
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestPilaDePilas(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[TDAPila.Pila[int]]()
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
	cola.Encolar(pila1)
	cola.Encolar(pila2)
	cola.Encolar(pila3)
	require.EqualValues(t, pila1, cola.VerPrimero())

	// Pruebo a desapilar e investigar sus elementos posteriormente.
	elemento := cola.Desencolar()
	require.EqualValues(t, 5, elemento.Desapilar())
	require.EqualValues(t, 4, elemento.Desapilar())
	require.EqualValues(t, 3, elemento.Desapilar())
	require.EqualValues(t, 2, elemento.Desapilar())
	require.EqualValues(t, 1, elemento.Desapilar())
	elemento = cola.Desencolar()
	require.EqualValues(t, 10, elemento.Desapilar())
	require.EqualValues(t, 9, elemento.Desapilar())
	require.EqualValues(t, 8, elemento.Desapilar())
	require.EqualValues(t, 7, elemento.Desapilar())
	require.EqualValues(t, 6, elemento.Desapilar())
	elemento = cola.Desencolar()
	require.EqualValues(t, 15, elemento.Desapilar())
	require.EqualValues(t, 14, elemento.Desapilar())
	require.EqualValues(t, 13, elemento.Desapilar())
	require.EqualValues(t, 12, elemento.Desapilar())
	require.EqualValues(t, 11, elemento.Desapilar())

	// Verifico que funcione como una pila vacia.
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia.", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())

}
