package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodoLista(dato, lista.primero)
	lista.primero = nuevoNodo
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodoLista(dato, nil)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	dato := lista.VerPrimero()
	lista.primero = lista.primero.siguiente
	lista.largo--
	if lista.largo == 0 {
		lista.ultimo = nil
	}

	return dato
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
	for nodo := lista.primero; nodo != nil; nodo = nodo.siguiente {
		if !visitar(nodo.dato) {
			break
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, anterior: nil, lista: lista}
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodoLista(dato, iter.actual)
	if iter.anterior == nil {
		iter.lista.primero = nuevoNodo
	} else {
		iter.anterior.siguiente = nuevoNodo
	}
	if iter.actual == nil {
		iter.lista.ultimo = nuevoNodo
	}
	iter.lista.largo++
	iter.actual = nuevoNodo
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	dato := iter.actual.dato
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}

	iter.lista.largo--
	if iter.lista.largo == 0 {
		iter.lista.ultimo = nil
	} else if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.siguiente

	return dato
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodoLista[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	nodoLista := new(nodoLista[T])
	nodoLista.dato = dato
	nodoLista.siguiente = siguiente

	return nodoLista
}
