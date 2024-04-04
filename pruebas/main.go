package main

import "fmt"

func Comparar(vector1 []int, vector2 []int) int {

	// En este punto es indiferente que arreglo usar como referencia de len ya que son del mismo tamaño.
	for i := 0; i < len(vector1) && i < len(vector2); i++ {

		if vector1[i] > vector2[i] {
			return 1
		} else if vector1[i] < vector2[i] {
			return -1
		}
	}

	// Si no tienen elementos mas grandes determino el mayor arreglo segun su tamaño.
	if len(vector1) > len(vector2) {
		return 1
	} else if len(vector1) < len(vector2) {
		return -1
	}

	// Serian iguales.
	return 0

}

func main() {
	a := []int{}
	b := []int{1, 2, 3}

	fmt.Println(Comparar(a, b))
}
