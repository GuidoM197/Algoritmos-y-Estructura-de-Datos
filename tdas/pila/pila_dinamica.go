package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	nuevaPila := new(pilaDinamica[T])
	datos := make([]T, 5)
	nuevaPila.datos = datos
	return nuevaPila
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == cap(p.datos) {
		nuevosDatos := make([]T, cap(p.datos)*2)
		copy(nuevosDatos, p.datos)
		p.datos = nuevosDatos
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad*4 == cap(p.datos) {
		nuevosDatos := make([]T, cap(p.datos)/2)
		copy(nuevosDatos, p.datos)
		p.datos = nuevosDatos
	}
	p.cantidad--
	return p.datos[p.cantidad]
}
