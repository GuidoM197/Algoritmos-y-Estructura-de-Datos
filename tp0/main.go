package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const rutaArch1 = "archivo1.in"
const rutaArch2 = "archivo2.in"

func recolectorDeDatos(ruta string) []int {

	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	lineas := bufio.NewScanner(archivo)

	sliceAux := []string{}

	for lineas.Scan() {
		sliceAux = append(sliceAux, lineas.Text())
	}

	sliceFinal := []int{}

	for _, valor := range sliceAux {
		numero, _ := strconv.Atoi(valor)
		sliceFinal = append(sliceFinal, numero)
	}

	return sliceFinal
}

func main() {

	arreglo1 := recolectorDeDatos(rutaArch1)
	arreglo2 := recolectorDeDatos(rutaArch2)

	mayorArreglo := ejercicios.Comparar(arreglo1, arreglo2)

	if mayorArreglo == 1 {
		ejercicios.Seleccion(arreglo1)

		for _, valor := range arreglo1 {
			fmt.Println(valor)
		}

	} else if mayorArreglo == -1 {
		ejercicios.Seleccion(arreglo2)

		for _, valor := range arreglo2 {
			fmt.Println(valor)
		}

	} else if mayorArreglo == 0 {
		ejercicios.Seleccion(arreglo1)
		fmt.Println("Son del identicos y se ven asi!.")

		for _, valor := range arreglo1 {
			fmt.Println(valor)
		}

	}

}
