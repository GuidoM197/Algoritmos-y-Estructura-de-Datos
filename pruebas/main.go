package main

import (
	"fmt"
	lista "tdas/lista"
)

func main() {
	lista_Enlazada := lista.CrearListaEnlazada[int]()
	res := []int{}

	for i := 0; i < 5; i++ {
		lista_Enlazada.InsertarUltimo(i)

	}
	lista_Enlazada.ToString()

	contador := 0
	lista_Enlazada.Iterar(func(i int) bool {
		if contador == 3 {
			return false
		}
		v := i * 2
		res = append(res, v)
		contador++
		return true
	})

	lista_Enlazada.ToString()
	fmt.Println(res)
}
