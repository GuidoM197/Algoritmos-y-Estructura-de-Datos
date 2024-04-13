package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tdas/cola"
	"tp1/operations"
)

func main() {

	file := bufio.NewScanner(os.Stdin)

	operationQueue := cola.CrearColaEnlazada[[]string]()

	for file.Scan() {
		sliceAux := strings.Split(file.Text(), " ")
		operationQueue.Encolar(sliceAux)

	}

	for !operationQueue.EstaVacia() {
		actualOperation := operationQueue.Desencolar()
		result, err := operations.IdentifyOperations(actualOperation)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	}
}
