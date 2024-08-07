package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero guarda un elemento tipo T pasado por parametro en la primera posición de la lista enlazada.
	InsertarPrimero(T)

	// InsertarUltimo guarda un elemento tipo T pasado por parametro en la ultima posición de la lista enlazada.
	InsertarUltimo(T)

	// BorrarPrimero elimina el elemento que se almacena en la primer posición de la lista enlazada.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento que se encuentre en la primer posición de la lista enlazada.
	VerPrimero() T

	// VerUltimo devuelve el elemento que se encuentre en la ultima posición de la lista enlazada.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista enlazada.
	Largo() int

	// Iterar recorre la lista enlazada hasta que se la función pasada por parametro devuelva false.
	Iterar(visitar func(T) bool)

	// Iterador crea un itrador externo el cual puede recorrer la lista y aplicar sus primitivas sobre ella.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual itera el arreglo y devuelve el elemento en el que esta actualmente parando.
	VerActual() T

	// HaySiguiente devuelve true en caso de que el elemento actual sea diferente de nil.
	HaySiguiente() bool

	// Siguiente se ubica en el siguiente elemento de la lista.
	Siguiente()

	// Insertar recibe un elemento tipo T por parametro y se encanrga de insertarlo en la lista.
	Insertar(T)

	// Borrar elimina de la lista el elemento indicado y lo devuelve.
	Borrar() T
}
