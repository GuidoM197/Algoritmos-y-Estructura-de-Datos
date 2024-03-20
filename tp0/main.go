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

	slice_aux := []string{}

	for lineas.Scan() {
		slice_aux = append(slice_aux, lineas.Text())
	}

	slice_final := []int{}

	for _, valor := range slice_aux {
		numero, _ := strconv.Atoi(valor)
		slice_final = append(slice_final, numero)
	}

	return slice_final
}

func main() {

	arreglo1 := recolectorDeDatos(rutaArch1)
	arreglo2 := recolectorDeDatos(rutaArch2)

	mayorArreglo := ejercicios.Comparar(arreglo1, arreglo2)

	if mayorArreglo == 1 {
		ejercicios.Seleccion(arreglo1)
		fmt.Println(arreglo1)

	} else if mayorArreglo == -1 {
		ejercicios.Seleccion(arreglo2)
		fmt.Println(arreglo2)
		
	} else if mayorArreglo == 0 {
		ejercicios.Seleccion(arreglo1)
		fmt.Println("Son del identicos y se ven asi!.", arreglo1)
	}

}
