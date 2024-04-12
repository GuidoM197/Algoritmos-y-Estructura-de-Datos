package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola"
	"tdas/pila"
)

const RUTA = "./prueba.txt"

func identificarNumeros(arreglo []string) (pila.Pila[int], []string) {
	nuevaPila := pila.CrearPilaDinamica[int]()
	arregloOperadores := []string{}
	for _, value := range arreglo {

		if value == "0" || value == "1" || value == "2" || value == "3" || value == "4" || value == "5" || value == "6" || value == "7" || value == "8" || value == "9" {
			number, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("No se pudo convertir correctamente: %v\n", err)
			}

			nuevaPila.Apilar(number)

		} else {
			arregloOperadores = append(arregloOperadores, value)
		}

	}
	return nuevaPila, arregloOperadores
}

func main() {
	archivo, err := os.Open(RUTA)
	if err != nil {
		fmt.Printf("No se pudo abrir el archivo correctamente: %v\n", err)
	}
	defer archivo.Close()

	lineas := bufio.NewScanner(archivo)

	colaOperaciones := cola.CrearColaEnlazada[[]string]()

	for lineas.Scan() {
		sliceAux := strings.Split(lineas.Text(), " ")
		colaOperaciones.Encolar(sliceAux)
	}

	for !colaOperaciones.EstaVacia() {
		operacionActual := colaOperaciones.Desencolar()
		pilaDeNumeros, operadores := identificarNumeros(operacionActual)

		fmt.Printf("%v ", operadores)
		pilaDeNumeros.VerPila()
	}
}
