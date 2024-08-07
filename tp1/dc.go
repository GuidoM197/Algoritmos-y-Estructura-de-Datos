package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1/operations"
)

func main() {
	//linea, _ := os.Open("./01_in")
	//defer linea.Close()

	file := bufio.NewScanner(os.Stdin)

	for file.Scan() {
		result, err := operations.IdentifyOperations(strings.Split(file.Text(), " "))
		if err != nil {
			fmt.Println("Entre")
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	}
}
