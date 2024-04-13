package main

import (
	"fmt"
	"tdas/cola"
	"tdas/pila"
)

func mismosOperadores(operadores cola.Cola[int], numeros pila.Pila[int]) bool {
	var (
		counterOperators int
		counterNumbers   int
		pilaAux          pila.Pila[int]
		colaAux          cola.Cola[int]
	)

	for !operadores.EstaVacia() && !numeros.EstaVacia() {

		colaAux.Encolar(operadores.Desencolar())
		pilaAux.Apilar(numeros.Desapilar())

		counterOperators++
		counterNumbers++
	}

	return counterNumbers == counterOperators
}

func main() {

	cola := cola.CrearColaEnlazada[int]()
	pila := pila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	pila.Apilar(1)
	pila.Apilar(1)
	pila.Apilar(1)

	cola.Encolar(1)
	cola.Encolar(1)
	cola.Encolar(1)

	fmt.Println(mismosOperadores(cola, pila))

	cola.VerCola()
	pila.VerPila()

}

/*
5 3 +
5 3 -
5 3 /
3 5 8 + +
3 5 8 + -
*/
