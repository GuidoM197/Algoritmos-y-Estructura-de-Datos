package guia_1

import (
	"fmt"
	TDAPila "tdas/pila"
)

// Guia de TDAS

func Ej1() {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := range 6 {
		pila.Apilar(i)
	}
	fmt.Println(pila.Multitope(3))
}
