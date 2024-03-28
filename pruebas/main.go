package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range arr {
		fmt.Println(arr[i])
	}
}
