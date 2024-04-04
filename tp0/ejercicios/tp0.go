package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}

	maximo, posicionMax := vector[0], 0

	for i := 0; i < len(vector); i++ {

		if vector[i] > maximo {
			posicionMax = i
			maximo = vector[i]
		}
	}

	return posicionMax

}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
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

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {

	for i := (len(vector) - 1); i > 0; i-- {
		posicionMax := Maximo(vector[:i+1])
		Swap(&vector[i], &vector[posicionMax])

	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[0] + Suma(vector[1:])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func EsCadenaCapicua(cadena string) bool {
	if 0 == len(cadena) || 1 == len(cadena) {
		return true
	}

	if cadena[0] != cadena[len(cadena)-1] {
		return false
	}

	return EsCadenaCapicua(cadena[1 : len(cadena)-1])
}
