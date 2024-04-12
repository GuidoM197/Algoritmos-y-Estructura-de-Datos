package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
<<<<<<< HEAD
	TDAPila "tdas/pila"
=======
	"tdas/cola"
	"tdas/pila"
>>>>>>> 26a25cf105ba5895fc716f82347d1f494ab78e94
)

const RUTA = "./prueba.txt"

<<<<<<< HEAD
func suma(pila TDAPila.Pila[int]) int {
	var res int
	var counter int
	for !pila.EstaVacia() && counter != 2 {
		res += pila.Desapilar()
		counter++
	}
	return res
}

func main() {
	//file := bufio.NewScanner(os.Stdin)
	pilaNumero := TDAPila.CrearPilaDinamica[int]()
	//pilaOperadores := TDAPila.CrearPilaDinamica[string]()

=======
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
>>>>>>> 26a25cf105ba5895fc716f82347d1f494ab78e94
	archivo, err := os.Open(RUTA)
	if err != nil {
		fmt.Printf("No se pudo abrir el archivo correctamente: %v\n", err)
	}
	defer archivo.Close()

	lineas := bufio.NewScanner(archivo)

<<<<<<< HEAD
	for lineas.Scan() {
		sliceAux := strings.Split(lineas.Text(), " ")

		centinela := true

		for centinela {

			var i int

			numero, err := strconv.Atoi(sliceAux[i])
			if err != nil {
				fmt.Printf("No se pudo convertir correctamente: %v\n", err)
			}

			pilaNumero.Apilar(numero)
			i++

			if sliceAux[i] == "+" || sliceAux[i] == "-" {
				centinela = false

			}
		}
		var i int
		if sliceAux[i] == "+" {
			fmt.Println(suma(pilaNumero))
		}

	}

=======
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
>>>>>>> 26a25cf105ba5895fc716f82347d1f494ab78e94
}
