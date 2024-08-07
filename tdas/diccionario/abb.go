package diccionario

import (
	"tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	valor     V
}

// funcCmp es un type que representa una funcion que compara dos elementos de tipo comparable,
// devuelve un entero que indica si el primer elemento es menor, igual o mayor
type funcCmp[K comparable] func(a, b K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcion_cmp
	return abb
}

func (a *abb[K, V]) Guardar(clave K, valor V) {
	if a.raiz == nil {
		a.raiz = crearNodo(clave, valor)
		a.cantidad++
	} else {
		a.guardarEnPosicion(clave, valor)
	}
}

func (a abb[K, V]) Pertenece(clave K) bool {
	if a.cantidad == 0 {
		return false
	}
	encontrado := a.buscar(clave, &a.raiz)
	if *encontrado != nil {
		return true
	}
	return false
}

func (a *abb[K, V]) Obtener(clave K) V {
	if a.cantidad == 0 {
		panic("La clave no pertenece al diccionario")
	}
	encontado := a.buscar(clave, &a.raiz)
	if *encontado == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*encontado).valor
}

func (a *abb[K, V]) Borrar(clave K) V {
	if a.raiz == nil {
		panic("La clave no pertenece al diccionario")

	}

	encontrado := a.buscar(clave, &a.raiz)
	if *encontrado == nil {
		panic("La clave no pertenece al diccionario")
	}

	valorDeEliminado := (*encontrado).valor

	if (*encontrado).izquierdo != nil && (*encontrado).derecho != nil {
		a.borrarDosHijos(encontrado)

	} else if unSoloHijo(encontrado) {
		a.borrarUnHijo(encontrado)

	} else {
		*encontrado = nil

	}
	a.cantidad--
	return valorDeEliminado
}

func (a abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) Iterar(f func(clave K, valor V) bool) {
	centinela := true
	a.iterarRecursivo(&a.raiz, f, &centinela)
}

func (a *abb[K, V]) IterarRango(desde, hasta *K, visitar func(clave K, valor V) bool) {
	centinela := true
	a.iterarRangoRecursivamente(a.raiz, desde, hasta, visitar, &centinela)
}

func (a *abb[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	pilaOrdenada := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if a.raiz != nil {
		pilaOrdenada = a.inicializarPila(a.raiz, desde, hasta, pilaOrdenada)
	}
	iterador := new(iteradorABB[K, V])
	iterador.pilaOrdenada = pilaOrdenada
	iterador.desde, iterador.hasta = desde, hasta
	iterador.arbol = a
	return iterador
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

type iteradorABB[K comparable, V any] struct {
	pilaOrdenada pila.Pila[*nodoAbb[K, V]]
	desde        *K
	hasta        *K
	arbol        *abb[K, V]
}

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	return !iter.pilaOrdenada.EstaVacia()
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pilaOrdenada.VerTope()
	return (*actual).clave, (*actual).valor
}

func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := iter.pilaOrdenada.Desapilar()
	if actual.derecho != nil {
		iter.pilaOrdenada = iter.arbol.actualizarPila(actual.derecho, iter.desde, iter.hasta, iter.pilaOrdenada)
	}
}

// ==================================== Funciones auxiliares ======================================= //

func (a *abb[K, V]) buscar(clave K, actual **nodoAbb[K, V]) **nodoAbb[K, V] {
	if actual == nil {
		return nil
	}
	rtaCmp := a.cmp(clave, (*actual).clave)

	if rtaCmp < 0 {
		if (*actual).izquierdo == nil {
			return &(*actual).izquierdo
		}
		return a.buscar(clave, &(*actual).izquierdo)
	} else if rtaCmp > 0 {
		if (*actual).derecho == nil {
			return &(*actual).derecho
		}
		return a.buscar(clave, &(*actual).derecho)
	}
	return actual
}

func (a *abb[K, V]) guardarEnPosicion(clave K, valor V) {
	encontrado := a.buscar(clave, &a.raiz)

	if (*encontrado) == nil {
		//si no hay nodo el actual creo un nodo y lo guardo en la posicion correspondiente
		nuevoNodo := crearNodo(clave, valor)
		*encontrado = nuevoNodo
		a.cantidad++

	} else {
		(*encontrado).valor = valor

	}
}

func (a *abb[K, V]) borrarDosHijos(actual **nodoAbb[K, V]) {
	padreReemplazo := *actual
	reemplazo := (*actual).izquierdo

	for reemplazo.derecho != nil {
		padreReemplazo = reemplazo
		reemplazo = reemplazo.derecho
	}

	(*actual).clave, (*actual).valor = reemplazo.clave, reemplazo.valor

	if padreReemplazo == *actual {
		padreReemplazo.izquierdo = reemplazo.izquierdo
	} else {
		padreReemplazo.derecho = reemplazo.izquierdo
	}
}

func (a *abb[K, V]) borrarUnHijo(actual **nodoAbb[K, V]) {
	if (*actual).izquierdo != nil {
		*actual = (*actual).izquierdo

	} else {
		*actual = (*actual).derecho
	}
}

func (a *abb[K, V]) iterarRecursivo(actual **nodoAbb[K, V], f func(clave K, valor V) bool, centinela *bool) {
	if *actual == nil || !*centinela {
		return
	}
	a.iterarRecursivo(&(*actual).izquierdo, f, centinela)

	if *centinela && !f((*actual).clave, (*actual).valor) {
		*centinela = false

	}
	if *centinela {
		a.iterarRecursivo(&(*actual).derecho, f, centinela)
	}
}

func (a *abb[K, V]) iterarRangoRecursivamente(actual *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, valor V) bool, centinela *bool) {
	if actual == nil || !*centinela {
		return
	}

	if desde != nil && a.cmp(actual.clave, *desde) < 0 {
		a.iterarRangoRecursivamente(actual.derecho, desde, hasta, visitar, centinela)

	} else if hasta != nil && a.cmp(actual.clave, *hasta) > 0 {
		a.iterarRangoRecursivamente(actual.izquierdo, desde, hasta, visitar, centinela)

	} else {
		a.iterarRangoRecursivamente(actual.izquierdo, desde, hasta, visitar, centinela)

		if *centinela && !visitar(actual.clave, actual.valor) {
			*centinela = false
		}

		if *centinela {
			a.iterarRangoRecursivamente(actual.derecho, desde, hasta, visitar, centinela)
		}
	}
}

func unSoloHijo[K comparable, V any](encontrado **nodoAbb[K, V]) bool {
	return ((*encontrado).izquierdo == nil && (*encontrado).derecho != nil) || ((*encontrado).izquierdo != nil && (*encontrado).derecho == nil)
}

func (a *abb[K, V]) actualizarPila(actual *nodoAbb[K, V], desde *K, hasta *K, pilaABB pila.Pila[*nodoAbb[K, V]]) pila.Pila[*nodoAbb[K, V]] {
	for actual != nil {
		// Si esta dentro del rango se apila el valor
		if (desde == nil || a.cmp((*actual).clave, *desde) >= 0) && (hasta == nil || a.cmp((*actual).clave, *hasta) <= 0) {
			pilaABB.Apilar(actual)
		}
		// Si es mayor cambiamos la referencia del nodo para el hijo izquierdo
		if desde == nil || a.cmp((*actual).clave, *desde) > 0 {
			actual = (*actual).izquierdo

		} else { // Sino cambiamos la referencia del nodo para para el hijo derecho
			actual = (*actual).derecho
		}
	}
	return pilaABB
}

func (a *abb[K, V]) inicializarPila(actual *nodoAbb[K, V], desde *K, hasta *K, pilaABB pila.Pila[*nodoAbb[K, V]]) pila.Pila[*nodoAbb[K, V]] {
	for actual != nil {
		// Si esta dentro del rango se apila el valor
		if (desde == nil || a.cmp((*actual).clave, *desde) >= 0) && (hasta == nil || a.cmp((*actual).clave, *hasta) <= 0) {
			pilaABB.Apilar(actual)
			if desde != nil && (*actual).clave == *desde {
				break
			}
		}
		// Si es mayor cambiamos la referencia del nodo para el hijo izquierdo
		if desde == nil || (a.cmp((*actual).clave, *desde) > 0 && (*actual).izquierdo != nil) {
			actual = (*actual).izquierdo
		} else if hasta == nil || (a.cmp((*actual).clave, *hasta) < 0 && (*actual).derecho != nil) { // Sino cambiamos la referencia del nodo para para el hijo derecho
			actual = (*actual).derecho
		} else {
			break
		}
	}
	return pilaABB
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.valor = valor
	return nodo
}
