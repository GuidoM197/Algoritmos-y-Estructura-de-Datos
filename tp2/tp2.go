package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tdas/diccionario"
	"tp2/operaciones"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	visitantes := diccionario.CrearHash[string, operaciones.DiccOrdenado]()
	urlsMasVisitadas := []operaciones.Par{}

	for s.Scan() {
		line := s.Text()
		input := strings.Fields(line)
		var err error
		if len(input) == 0 {
			continue
		}
		comando := input[0]

		if comando == "agregar_archivo" {
			if len(input) < 2 {
				err = fmt.Errorf("Error en comando agregar_archivo")
			} else {
				err = operaciones.AgregarArchivo(input[1], visitantes, &urlsMasVisitadas)
			}

		} else if comando == "ver_mas_visitados" {
			if len(input) < 2 {
				err = fmt.Errorf("Error en comando ver_mas_visitados")
			} else {
				operaciones.VerMasVisitados(input[1], urlsMasVisitadas)
			}

		} else if comando == "ver_visitantes" {
			if len(input) < 3 {
				err = fmt.Errorf("Error en comando ver_visitantes")
			} else {
				operaciones.VerVisitantes(visitantes, input[1], input[2])
			}
		} else {
			err = fmt.Errorf("Error en comando %s", comando)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
