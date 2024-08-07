package cola_prioridad

import "fmt"

const (
	_TAMANIO_INICIAL                 = 10
	_CAPACIDAD_INICIAL               = 0
	_AUMENTAR_CAPACIDAD              = 2
	_DISMINUIR_CAPACIDAD             = 2
	_PARAMETRO_CONDICION_REDIMENCION = 4
)

type heap[T comparable] struct {
	cola     []T
	cantidad int
	cmp      func(a, b T) int
}

func CrearHeap[T comparable](cmp func(a, b T) int) ColaPrioridad[T] {
	return &heap[T]{cola: make([]T, _TAMANIO_INICIAL), cantidad: _CAPACIDAD_INICIAL, cmp: cmp}
}

func CrearHeapArr[T comparable](arreglo []T, cmp func(a, b T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap(cmp)
	}
	aux := make([]T, len(arreglo))
	copy(aux, arreglo)
	heap := &heap[T]{cola: aux, cantidad: len(arreglo), cmp: cmp}
	heapify(heap.cola, cmp)
	return heap
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(valor T) {
	if h.cantidad == len(h.cola) {
		h.cola = redimension(cap(h.cola)*_AUMENTAR_CAPACIDAD, h.cola)
	}
	h.cola[h.cantidad] = valor
	h.cantidad++
	upHeap((h.cantidad - 1), h.cola, h.cmp)
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.cola[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	if h.cantidad*_PARAMETRO_CONDICION_REDIMENCION == cap(h.cola) {
		h.cola = redimension(cap(h.cola)/_DISMINUIR_CAPACIDAD, h.cola)
	}
	valor := h.VerMax()
	h.cantidad--
	h.cola[0], h.cola = h.cola[h.cantidad], h.cola[:h.cantidad]
	downHeap(0, h.cola, h.cmp)
	return valor
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

// ----------------------- Funciones auxiliares ----------------------- //

func redimension[T any](nuevaCapacidad int, datos []T) []T {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, datos)
	return nuevosDatos
}

func upHeap[T any](posicion int, elementos []T, cmp func(a, b T) int) {
	padre := (posicion - 1) / 2
	if cmp(elementos[padre], elementos[posicion]) < 0 {
		swap(&elementos[padre], &elementos[posicion])
		upHeap(padre, elementos, cmp)
	}
}

func downHeap[T any](posicion int, elementos []T, cmp func(a, b T) int) {
	if posicion > len(elementos) {
		return
	}
	var (
		hijoIzquierdo = posicion*2 + 1
		hijoDerecho   = posicion*2 + 2
		mayor         = posicion
	)

	if hijoIzquierdo < len(elementos) && cmp(elementos[mayor], elementos[hijoIzquierdo]) < 0 {
		mayor = hijoIzquierdo
	}

	if hijoDerecho < len(elementos) && cmp(elementos[mayor], elementos[hijoDerecho]) < 0 {
		mayor = hijoDerecho
	}

	if mayor != posicion {
		swap(&elementos[posicion], &elementos[mayor])
		downHeap(mayor, elementos, cmp)
	}
}

func HeapSort[T any](elementos []T, cmp func(a, b T) int) {
	heapify(elementos, cmp)
	cantidad := len(elementos) - 1
	for i := len(elementos) - 1; i >= 0; i-- {
		swap(&elementos[i], &elementos[0])
		downHeap(0, elementos[:cantidad], cmp)
		cantidad--
	}
}

func swap[T any](x, y *T) {
	*x, *y = *y, *x
}

func heapify[T any](arreglo []T, cmp func(a, b T) int) {
	for i := len(arreglo)/2 - 1; i >= 0; i-- {
		downHeap(i, arreglo, cmp)
	}
}

func (h *heap[T]) ToString() {
	fmt.Printf("<| Primero |")
	for !h.EstaVacia() {
		fmt.Printf(" <- %v", h.VerMax())
		h.Desencolar()
	}
	fmt.Printf(" <| Fin |\n")
}
