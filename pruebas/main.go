package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

const RUTA = "./prueba.txt"

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

	archivo, err := os.Open(RUTA)
	if err != nil {
		fmt.Printf("No se pudo abrir el archivo correctamente: %v\n", err)
	}
	defer archivo.Close()

	lineas := bufio.NewScanner(archivo)

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

}
