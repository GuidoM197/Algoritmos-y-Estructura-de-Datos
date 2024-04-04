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

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("No se pudo abrir el archivo correctamente: %v\n", err)
	}
	defer archivo.Close()

	lineas := bufio.NewScanner(archivo)

	sliceAux := []string{}

	for lineas.Scan() {
		sliceAux = append(sliceAux, lineas.Text())
	}

	sliceFinal := []int{}

	for _, valor := range sliceAux {
		numero, err := strconv.Atoi(valor)
		if err != nil {
			fmt.Printf("No se pudo convertir correctamente: %v\n", err)
		}
		sliceFinal = append(sliceFinal, numero)
	}

	return sliceFinal
}

func main() {

	var resultado []int

	arreglo1 := recolectorDeDatos(rutaArch1)
	arreglo2 := recolectorDeDatos(rutaArch2)

	mayorArreglo := ejercicios.Comparar(arreglo1, arreglo2)

	if mayorArreglo == 1 || mayorArreglo == 0 {
		ejercicios.Seleccion(arreglo1)
		resultado = arreglo1

	} else if mayorArreglo == -1 {
		ejercicios.Seleccion(arreglo2)
		resultado = arreglo2

	}

	for _, valor := range resultado {
		fmt.Println(valor)
	}

}
