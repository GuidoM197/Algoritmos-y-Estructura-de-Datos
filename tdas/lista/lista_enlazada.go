package lista

import "fmt"

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevoNodo := crearNodo[T](valor, nil)
	lista.largo++

	if lista.primero == nil {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo

	} else {
		actual := lista.primero
		lista.primero = nuevoNodo
		lista.primero.prox = actual

	}
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevoNodo := crearNodo[T](valor, nil)
	lista.largo++

	if lista.ultimo == nil {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo

	} else {
		lista.ultimo.prox = nuevoNodo
		lista.ultimo = lista.ultimo.prox

	}
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	valor := lista.primero.dato
	lista.primero = lista.primero.prox
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return valor
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil && visitar(actual.dato) {
		actual = actual.prox
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista.primero, nil, lista}
}

//////////////////////////////////////////////////////////////////////////////////

func (lista *listaEnlazada[T]) ToString() {
	actual := lista.primero
	fmt.Printf("|> Primero |")
	for actual != nil {
		fmt.Printf(" -> %v", actual.dato)
		actual = actual.prox
	}
	fmt.Printf(" | Fin |>\n")
}

/////////////////////////////////////////////////////////////////////////////////

type iterListaEnlazada[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if iterador.lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.prox
}

func (iterador *iterListaEnlazada[T]) Insertar(valor T) {
	// Lista vacia.
	if iterador.actual == nil {
		iterador.actual = crearNodo(valor, nil)
		iterador.anterior.prox = iterador.actual
		iterador.lista.ultimo = iterador.actual

		// Primer elemento, no vacia.
	} else if iterador.anterior == nil {
		iterador.actual = crearNodo(valor, iterador.actual)
		iterador.lista.primero = iterador.actual

		// Ultimo elemento.
	} else if iterador.actual.prox == nil {
		iterador.actual.prox = crearNodo(valor, nil)
		iterador.lista.ultimo = iterador.actual.prox

		// Elemento intermedio.
	} else {
		iterador.anterior.prox = crearNodo(valor, iterador.actual)
		iterador.actual = iterador.anterior.prox

	}
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if iterador.actual == nil {
		panic("El iterador termino de iterar")
	}
	iterador.lista.largo--

	if iterador.anterior == nil {
		dato := iterador.actual.dato
		iterador.actual = iterador.actual.prox
		if iterador.lista.ultimo == iterador.lista.primero {
			iterador.lista.ultimo = nil

		}
		iterador.lista.primero = iterador.lista.primero.prox
		return dato
	}
	dato := iterador.actual.dato
	iterador.anterior.prox = iterador.actual.prox
	if iterador.anterior.prox == nil {
		iterador.lista.ultimo = iterador.anterior

	}
	return dato
}

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

func crearNodo[T any](valor T, prox *nodo[T]) *nodo[T] {
	return &nodo[T]{valor, prox}
}
