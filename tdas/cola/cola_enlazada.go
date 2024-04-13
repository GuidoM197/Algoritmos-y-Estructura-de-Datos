package cola

import "fmt"

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.valor
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := crearNodo(valor, nil)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	valor := cola.primero.valor
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return valor
}

type nodo[T any] struct {
	valor T
	prox  *nodo[T]
}

func crearNodo[T any](dato T, proximo *nodo[T]) *nodo[T] {
	return &nodo[T]{dato, proximo}
}

func (cola colaEnlazada[T]) VerCola() {
	fmt.Printf("<| Primero |")
	for !cola.EstaVacia() {
		fmt.Printf(" <- %v", cola.VerPrimero())
		cola.Desencolar()
	}
	fmt.Printf(" <| Fin |\n")
}
