package main

import (
	"fmt"
)

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := crearNodoCola(valor, nil)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia.")
	}
	return cola.primero.valor
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia.")
	}
	valor := cola.primero.valor
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return valor
}

func (cola *colaEnlazada[int]) VerCola() {
	fmt.Printf("Frente <| ")
	for cola.primero != nil {
		fmt.Printf("%v", cola.primero.valor)
		cola.Desencolar()
		if cola.primero != nil {
			fmt.Printf(" <- ")
		}
	}
	fmt.Printf(" <| Fondo")
}

type nodoCola[T any] struct {
	valor T
	prox  *nodoCola[T]
}

func crearNodoCola[T any](dato T, proximo *nodoCola[T]) *nodoCola[T] {
	return &nodoCola[T]{valor: dato, prox: proximo}
}

func main() {
	//	a := []int{}
	//	b := []int{1, 2, 3}

	cola := CrearColaEnlazada[int]()

	cola.Encolar(1)

	cola.Encolar(2)

	cola.Encolar(3)

	cola.Encolar(4)

	fmt.Println(cola.VerPrimero())
	cola.VerCola()

}

type Cola[T any] interface {
	EstaVacia() bool
	Encolar(T)
	Desencolar() T
	VerCola()
	VerPrimero() T
}
