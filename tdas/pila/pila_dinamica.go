package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func redimension[T any](cantidad, capacidad int, datos []T) []T {
	if cantidad == capacidad {
		nuevosDatos := make([]T, cap(datos)*2)
		copy(nuevosDatos, datos)
		return nuevosDatos

	} else if cantidad*4 == capacidad {
		nuevosDatos := make([]T, cap(datos)/2)
		copy(nuevosDatos, datos)
		return nuevosDatos

	}
	return datos
}

func CrearPilaDinamica[T any]() Pila[T] {
	nuevaPila := new(pilaDinamica[T])
	datos := make([]T, 5)
	nuevaPila.datos = datos
	return nuevaPila
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
	pila.datos = redimension(pila.cantidad, cap(pila.datos), pila.datos)
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.datos = redimension(pila.cantidad, cap(pila.datos), pila.datos)
	pila.cantidad--
	return pila.datos[pila.cantidad]
}

// Func de la guia.

func (pila pilaDinamica[T]) Multitope(n int) []T {
	cantidad := min(pila.cantidad, n)
	topes := make([]T, cantidad)
	for i := 0; i < cantidad; i++ {
		topes[i] = pila.datos[pila.cantidad-i-1]
	}
	return topes
}
