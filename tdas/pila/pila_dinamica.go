package pila

import "fmt"

/* Definición del struct pila proporcionado por la cátedra. */

const CANTIDAD_INICIAL = 0
const CAPACIDAD_INICIAL = 5
const AUMENTAR_CAPACIDAD = 2
const DISMINUIR_CAPACIDAD = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, CAPACIDAD_INICIAL), cantidad: CANTIDAD_INICIAL}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		pila.datos = redimension(cap(pila.datos)*DISMINUIR_CAPACIDAD, pila.datos)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad*4 == cap(pila.datos) {
		pila.datos = redimension(cap(pila.datos)/DISMINUIR_CAPACIDAD, pila.datos)
	}
	pila.cantidad--
	return pila.datos[pila.cantidad]
}

func redimension[T any](nuevaCapacidad int, datos []T) []T {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, datos)
	return nuevosDatos
}

func (pila *pilaDinamica[T]) VerPila() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	fmt.Printf("<|fin ")
	for i := 0; i < pila.cantidad; i++ {
		fmt.Printf("<- %v ", pila.datos[i])
	}
	fmt.Printf("<|tope")
}
