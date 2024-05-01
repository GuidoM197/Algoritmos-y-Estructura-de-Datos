package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1/operations"
)

func main() {
	file := bufio.NewScanner(os.Stdin)

	for file.Scan() {
		result, err := operations.IdentifyOperations(strings.Split(file.Text(), " "))
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	}
}
