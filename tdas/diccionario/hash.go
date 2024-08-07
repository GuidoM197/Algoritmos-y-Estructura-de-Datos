package diccionario

import (
	"fmt"
	"tdas/lista"
)

const (
	TAMANIOINICIAL         = 11
	CANTIDADINICIAL        = 0
	FACTORDECARGAAUMENTAR  = 2
	FACTORDECARGADISMINUIR = 1
	AUMENTARREDIMENCION    = 2
	DISMINUIRREDIMENCION   = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []lista.Lista[parClaveValor[K, V]]
	tamanio  int
	cantidad int
}

func (h *hashAbierto[K, V]) crearTabla(tamanio int) {
	h.tabla = make([]lista.Lista[parClaveValor[K, V]], tamanio)
	h.tamanio = tamanio
	h.cantidad = CANTIDADINICIAL
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashAbierto[K, V])
	hash.crearTabla(TAMANIOINICIAL)
	return hash
}

func (h *hashAbierto[K, V]) Guardar(clave K, valor V) {
	factorDeCarga := h.cantidad / h.tamanio
	if factorDeCarga > FACTORDECARGAAUMENTAR {
		h.redimension(h.tamanio * AUMENTARREDIMENCION)
	}
	h.insertarElemento(clave, valor)
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	pos := conseguirPosicion(clave, h.tamanio)

	// Verifico que la pos de la clave prtenesca a una lista.
	if h.tabla[pos] == nil {
		return false
	}

	// Busco en que lugar de mi lista esta la clave.
	buscado := busqueda(h.tabla[pos], clave)

	// Devuelvo si la clave es igual a la clave buscada.
	return buscado.HaySiguiente()
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	if h.Cantidad() == 0 {
		panic("La clave no pertenece al diccionario")

	}

	pos := conseguirPosicion(clave, h.tamanio)
	buscado := busqueda(h.tabla[pos], clave)
	if !buscado.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}

	return buscado.VerActual().valor
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	if h.Cantidad() == 0 {
		panic("La clave no pertenece al diccionario")

	}

	factorDeCarga := h.cantidad / h.tamanio
	if factorDeCarga < FACTORDECARGADISMINUIR {
		h.redimension(h.tamanio / DISMINUIRREDIMENCION)
	}

	valor := h.borrarElemento(clave)
	return valor
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h hashAbierto[K, V]) Iterar(vistar func(clave K, dato V) bool) {
	condicionDeCorte := true
	for _, list := range h.tabla {
		if list != nil && condicionDeCorte {
			for iter := list.Iterador(); iter.HaySiguiente(); iter.Siguiente() {

				if !vistar(iter.VerActual().clave, iter.VerActual().valor) {
					condicionDeCorte = false
					break
				}
			}
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	for i := range h.tabla {
		if h.tabla[i] != nil {
			return &iterDicc[K, V]{i, h, h.tabla[i].Iterador()}
		}
	}
	return &iterDicc[K, V]{h.tamanio, h, nil}

}

// ----------------- Funciónes Iterador Externo --------------------- //

type iterDicc[K comparable, V any] struct {
	actual        int
	tablaActual   *hashAbierto[K, V]
	iteradorTabla lista.IteradorLista[parClaveValor[K, V]]
}

func (iter *iterDicc[K, V]) HaySiguiente() bool {
	return iter.actual < iter.tablaActual.tamanio
}

func (iter *iterDicc[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	elemento := iter.iteradorTabla.VerActual()
	return elemento.clave, elemento.valor
}

func (iter *iterDicc[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	if iter.iteradorTabla.HaySiguiente() {
		iter.iteradorTabla.Siguiente()
	}

	iter.actualizarIteradorActual()
}

// -------------------- Funciónes Auxiliares ------------------------ //

func (iter *iterDicc[K, V]) actualizarIteradorActual() {
	if !iter.iteradorTabla.HaySiguiente() {
		iter.actual++
		for iter.HaySiguiente() && iter.tablaActual.tabla[iter.actual] == nil {
			iter.actual++

		}
		if iter.HaySiguiente() {
			iter.iteradorTabla = iter.tablaActual.tabla[iter.actual].Iterador()

		}
	}
}

func (h *hashAbierto[K, V]) redimension(nuevoTamanio int) {
	nuevoHash := new(hashAbierto[K, V])
	nuevoHash.crearTabla(nuevoTamanio)

	for _, listaActual := range h.tabla {
		if listaActual != nil {
			for iter := listaActual.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
				nuevoHash.insertarElemento(iter.VerActual().clave, iter.VerActual().valor)

			}
		}
	}
	h.tabla, h.tamanio = nuevoHash.tabla, nuevoTamanio
}

func (h *hashAbierto[K, V]) insertarElemento(clave K, valor V) {
	pos := conseguirPosicion(clave, h.tamanio)

	// Verifico si la lista en esa pos esta vacia, si lo esta inicializo una lista y agrego al final
	if h.tabla[pos] == nil {
		nuevaLista := lista.CrearListaEnlazada[parClaveValor[K, V]]()
		h.tabla[pos] = nuevaLista

	} else {
		buscado := busqueda(h.tabla[pos], clave)
		if buscado.HaySiguiente() {
			buscado.Borrar()
			h.cantidad--

		}
	}

	h.tabla[pos].InsertarUltimo(parClaveValor[K, V]{clave: clave, valor: valor})
	h.cantidad++
}

func (h *hashAbierto[K, V]) borrarElemento(clave K) V {
	pos := conseguirPosicion(clave, h.tamanio)

	buscado := busqueda(h.tabla[pos], clave)
	if !buscado.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}

	valor := buscado.VerActual().valor
	buscado.Borrar()

	//si la lista en la que borre queda vacia la seteo la csailla a nil
	if h.tabla[pos].EstaVacia() {
		h.tabla[pos] = nil

	}
	h.cantidad--
	return valor
}

// conseguirPosicion se le pasa una clave y el tamaño del arreglo y devuelve la clave hasheada.
func conseguirPosicion[K comparable](clave K, tamanio int) int {
	data := convertirABytes(clave)
	claveHash := xxhash64(data)

	return int(claveHash) % tamanio
}

/*
busqueda se le pasa por parámetro una clave y la posición del arreglo donde se desea buscar la misma,
en caso de que esta exista, se devuelve un itrador en la posición donde se encuentra,
en caso contrario, un iterador al final de la lista.
*/
func busqueda[K comparable, V any](listaActual lista.Lista[parClaveValor[K, V]], clave K) lista.IteradorLista[parClaveValor[K, V]] {
	if listaActual == nil {
		panic("La clave no pertenece al diccionario")
	}
	iter := listaActual.Iterador()

	for iter = listaActual.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			return iter

		}
	}

	// Si llega a este punto es porque no existe, entonces devuelve el iterador en la ultima pos.
	return iter
}

// toString nos sirve para observar el comportamiento del diccionario.
func (h *hashAbierto[K, V]) toString() {
	for i, listaActual := range h.tabla {
		fmt.Printf("Pos Hash: %d ->", i)
		if listaActual != nil {
			for iter := listaActual.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
				fmt.Printf(" %v: %v ->", iter.VerActual().clave, iter.VerActual().valor)
			}
		}
		fmt.Println(" nil")
	}
}
